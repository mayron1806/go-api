package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mayron1806/go-api/internal/helper"
	"github.com/mayron1806/go-api/internal/model"
	"gorm.io/gorm"
)

type ResetPasswordRequest struct {
	Token    string `json:"token" validate:"required"`
	Password string `json:"password" validate:"required,gte=6,lte=50"`
}

func (h *AuthHandler) ResetPassword(c *gin.Context) {
	var request ResetPasswordRequest
	if !h.ValidateRequest(c, &request) {
		return
	}

	var token model.Token
	err := h.db.Where("key = ?", request.Token).Joins("User").First(&token).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			h.ResponseError(c, http.StatusBadRequest, "token not found")
		} else {
			h.ResponseError(c, http.StatusBadRequest, "error finding token: %s", err.Error())
		}
		return
	}

	if token.ExpiresAt.Before(time.Now()) {
		h.ResponseError(c, http.StatusBadRequest, "invalid token or expired")
		return
	}
	if token.Type != model.ResetPassword {
		h.ResponseError(c, http.StatusBadRequest, "invalid token or expired")
		return
	}

	hashedPassword, err := helper.HashPassword(request.Password)
	if err != nil {
		h.ResponseError(c, http.StatusBadRequest, "error hashing password: %s", err.Error())
		return
	}

	token.User.Password = hashedPassword
	token.User.Challenge = model.UserChallengeNone
	err = h.db.Save(&token.User).Error
	if err != nil {
		h.ResponseError(c, http.StatusBadRequest, "error saving user: %s", err.Error())
		return
	}
	err = h.db.Delete(&token).Error
	if err != nil {
		h.ResponseError(c, http.StatusBadRequest, "error deleting token: %s", err.Error())
		return
	}
	c.Status(http.StatusOK)
}
