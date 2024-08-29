package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mayron1806/go-api/config"
	"github.com/mayron1806/go-api/internal/helper"
	"github.com/mayron1806/go-api/internal/model"
	"github.com/mayron1806/go-api/internal/template"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Account  string `json:"account" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (h *AuthHandler) Login(c *gin.Context) {
	var request LoginRequest
	if !h.ValidateRequest(c, &request) {
		return
	}

	var user model.User
	if helper.VerifyIsEmail(request.Account) {
		if err := h.db.Where("email = ?", request.Account).First(&user).Error; err != nil {
			h.ResponseError(c, http.StatusBadRequest, "error finding user: %s", err.Error())
			return
		}
	} else {
		if err := h.db.Where("name = ?", request.Account).First(&user).Error; err != nil {
			h.ResponseError(c, http.StatusBadRequest, "error finding user: %s", err.Error())
			return
		}
	}

	if !helper.CheckPasswordHash(request.Password, user.Password) {
		h.ResponseError(c, http.StatusBadRequest, "invalid credentials")
		return
	}
	if user.Challenge == model.UserChallengeVerifyEmail {
		// send email to user
		var token model.Token
		err := h.db.Where("user_id = ? AND type = ?", user.ID, model.ActiveAccount).First(&token).Error
		if err != nil {
			if err != gorm.ErrRecordNotFound {
				h.ResponseError(c, http.StatusBadRequest, "account not activated, error finding token: %s", err.Error())
				return
			}
			token = model.Token{
				Key:       uuid.New(),
				UserID:    user.ID,
				Type:      model.ActiveAccount,
				ExpiresAt: time.Now().Add(time.Hour * 24),
			}
			err = h.db.Create(&token).Error
			if err != nil {
				h.ResponseError(c, http.StatusBadRequest, "account not activated, error creating token: %s", err.Error())
			}
		}
		h.emailService.SendEmail(user.Email, "Account not activated", template.GetActiveAccountTemplate(token.Key.String()))
		h.ResponseError(c, http.StatusBadRequest, "user not active")
		return
	}

	permissions, err := h.queryUser.GetUserPermissions(user.ID)
	if err != nil {
		h.ResponseError(c, http.StatusBadRequest, "login error: %s", err.Error())
		return
	}
	// generate tokens
	accessToken, accessTokenError := h.jwtService.GenerateAccessToken(user, "credentials", permissions)
	if accessTokenError != nil {
		h.ResponseError(c, http.StatusBadRequest, "login error: %s", accessTokenError.Error())
		return
	}
	refreshToken := model.Token{
		Key:       uuid.New(),
		UserID:    user.ID,
		Type:      model.RefreshToken,
		Payload:   nil,
		ExpiresAt: time.Now().Add(time.Duration(config.GetEnv().JWT_REFRESH_TOKEN_DURATION)),
	}
	h.db.Create(&refreshToken)
	h.SetCookie(c, "access-token", accessToken.Token, int(accessToken.ExpiresAt.Sub(time.Now()).Seconds()))
	h.SetCookie(c, "expires-at", accessToken.ExpiresAt.Format(time.RFC3339), int(accessToken.ExpiresAt.Sub(time.Now()).Seconds()))
	h.SetCookie(c, "refresh-token", refreshToken.Key.String(), int(refreshToken.ExpiresAt.Sub(time.Now()).Seconds()))
	c.JSON(http.StatusOK, gin.H{"accessToken": accessToken.Token, "refreshToken": refreshToken.Key.String(), "expiresAt": accessToken.ExpiresAt})
}
