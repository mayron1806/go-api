package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mayron1806/go-api/internal/model"
)

type ActiveAccountRequest struct {
	Key string `json:"key" validate:"required"`
}

func (h *AuthHandler) ActiveAccount(c *gin.Context) {
	var request ActiveAccountRequest
	if !h.ValidateRequest(c, &request) {
		return
	}

	var token model.Token
	result := h.db.Where("key = ? AND type = ?", request.Key, model.ActiveAccount).First(&token)
	if result.Error != nil {
		h.ResponseError(c, http.StatusBadRequest, "error finding token: %s", result.Error.Error())
		return
	}
	if token.ExpiresAt.Before(time.Now()) {
		h.ResponseError(c, http.StatusBadRequest, "token expired")
		return
	}

	user := model.User{}
	err := h.db.Where("id = ?", token.UserID).First(&user).Error
	if err != nil {
		h.ResponseError(c, http.StatusBadRequest, "error finding user: %s", err.Error())
		return
	}

	user.Challenge = model.UserChallengeNone

	err = h.db.Save(&user).Error
	if err != nil {
		h.ResponseError(c, http.StatusBadRequest, "error saving user: %s", err.Error())
		return
	}
	err = h.db.Delete(&token).Error
	if err != nil {
		h.ResponseError(c, http.StatusBadRequest, "error deleting token: %s", err.Error())
		return
	}
	if result.Error != nil {
		h.ResponseError(c, http.StatusBadRequest, "error updating user: %s", result.Error.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}
