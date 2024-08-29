package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mayron1806/go-api/config"
	"github.com/mayron1806/go-api/internal/goauth2"
	"github.com/mayron1806/go-api/internal/model"
)

func (h *AuthHandler) OAuthCallback(c *gin.Context) {
	provider := c.Param("provider")
	authorizedToken, err := goauth2.Authorize(provider, c.Request.URL.Query())
	if err != nil {
		h.ResponseError(c, http.StatusBadRequest, "error authorizing: %s", err.Error())
		return
	}

	var userWithEmail model.User
	err = h.db.
		Joins("left join social_providers on social_providers.user_id = users.id").
		Attrs(model.User{
			Name:      authorizedToken.Name,
			Email:     authorizedToken.Email,
			Avatar:    authorizedToken.Avatar,
			Challenge: model.UserChallengeNone,
		}).
		FirstOrCreate(&userWithEmail, "users.email = ?", authorizedToken.Email).
		Error
	if err != nil {
		h.ResponseError(c, http.StatusBadRequest, "error authorizing: %s", err.Error())
		return
	}

	// add provider if not found
	var socialProvider model.SocialProvider
	err = h.db.
		Where("provider = ? AND user_id = ?", provider, userWithEmail.ID).
		Attrs(model.SocialProvider{
			Email:         authorizedToken.Email,
			EmailVerified: true,
			Active:        true,
			Provider:      provider,
			ProviderID:    authorizedToken.ProviderID,
			Avatar:        authorizedToken.Avatar,
			UserID:        userWithEmail.ID,
		}).
		FirstOrCreate(&socialProvider).Error

	if err != nil {
		h.ResponseError(c, http.StatusBadRequest, "error authorizing: %s", err.Error())
		return
	}

	permissions, err := h.queryUser.GetUserPermissions(userWithEmail.ID)
	if err != nil {
		h.ResponseError(c, http.StatusBadRequest, "error authorizing: %s", err.Error())
		return
	}
	// generate tokens
	accessToken, accessTokenError := h.jwtService.GenerateAccessToken(userWithEmail, socialProvider.Provider, permissions)
	if accessTokenError != nil {
		h.ResponseError(c, http.StatusBadRequest, "login error: %s", accessTokenError.Error())
		return
	}
	refreshToken := model.Token{
		Key:       uuid.New(),
		UserID:    userWithEmail.ID,
		Type:      model.RefreshToken,
		Payload:   authorizedToken,
		ExpiresAt: time.Now().Add(time.Duration(config.GetEnv().JWT_REFRESH_TOKEN_DURATION)),
	}
	h.db.Create(&refreshToken)
	h.SetCookie(c, "access-token", accessToken.Token, int(accessToken.ExpiresAt.Sub(time.Now()).Seconds()))
	h.SetCookie(c, "expires-at", accessToken.ExpiresAt.Format(time.RFC3339), int(accessToken.ExpiresAt.Sub(time.Now()).Seconds()))
	h.SetCookie(c, "refresh-token", refreshToken.Key.String(), int(refreshToken.ExpiresAt.Sub(time.Now()).Seconds()))
	c.JSON(http.StatusOK, gin.H{"accessToken": accessToken.Token, "refreshToken": refreshToken.Key.String(), "expiresAt": accessToken.ExpiresAt})
}
