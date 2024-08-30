package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	tokens, err := h.authService.GenerateTokens(
		&userWithEmail,
		socialProvider.Provider,
		permissions,
		&model.RefreshTokenPayload{Type: socialProvider.Provider, Oauth: *authorizedToken},
	)

	h.SetTokenCookies(c, tokens)
	c.JSON(http.StatusOK, gin.H{"accessToken": tokens.AccessToken.Token, "refreshToken": tokens.RefreshToken, "expiresAt": tokens.AccessToken.ExpiresAt})
}
