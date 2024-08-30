package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mayron1806/go-api/internal/goauth2"
	"github.com/mayron1806/go-api/internal/model"
)

type RefreshTokenRequest struct {
	Token string `json:"token" validate:"required"`
}

func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var request RefreshTokenRequest
	if !h.ValidateRequest(c, &request) {
		return
	}

	var token model.Token
	err := h.db.Where("key = ?", request.Token).Joins("User").First(&token).Error
	if err != nil {
		h.ResponseError(c, http.StatusBadRequest, "token not found: %s", err.Error())
		return
	}
	if token.ExpiresAt.Before(time.Now()) {
		h.db.Where("key = ?", token.Key).Delete(&model.Token{})
		h.ResponseError(c, http.StatusBadRequest, "token expired")
		return
	}
	if token.Type != model.RefreshToken {
		h.ResponseError(c, http.StatusBadRequest, "invalid token type")
		return
	}
	if token.Payload == nil {
		h.ResponseError(c, http.StatusBadRequest, "invalid token content")
		return
	}

	permissions, err := h.queryUser.GetUserPermissions(token.UserID)
	if err != nil {
		h.ResponseError(c, http.StatusBadRequest, "error getting user permissions: %s", err.Error())
		return
	}

	payloadMap, ok := token.Payload.(map[string]interface{})
	if !ok {
		h.ResponseError(c, http.StatusBadRequest, "invalid token content")
		return
	}
	payloadBytes, err := json.Marshal(payloadMap)
	if err != nil {
		h.ResponseError(c, http.StatusBadRequest, "invalid token content")
		return
	}
	var payload model.RefreshTokenPayload
	err = json.Unmarshal([]byte(payloadBytes), &payload)
	if err != nil {
		h.ResponseError(c, http.StatusBadRequest, "invalid token content")
		return
	}
	var OAuth *goauth2.AuthToken
	if payload.Type != "credentials" {
		h.Logger.Infof("revalidate oauth token: %s", payload.Oauth.AccessToken)
		OAuth, err = goauth2.RevalidateToken(payload.Type, payload.Oauth.AccessToken)
		if err != nil {
			h.ResponseError(c, http.StatusBadRequest, "token authentication error: %s", err.Error())
			return
		}
	}
	tx := h.db.Begin()
	tx.Where("key = ?", token.Key).Delete(&model.Token{})
	tokens, err := h.authService.GenerateTokens(
		&token.User,
		payload.Type,
		permissions,
		&model.RefreshTokenPayload{Type: payload.Type, Oauth: *OAuth},
	)
	if err != nil {
		tx.Rollback()
		h.ResponseError(c, http.StatusBadRequest, "login error: %s", err.Error())
		return
	}
	tx.Commit()
	h.SetTokenCookies(c, tokens)
	c.JSON(http.StatusOK, tokens)
	return
}
