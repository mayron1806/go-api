package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mayron1806/go-api/internal/model"
	"github.com/mayron1806/go-api/internal/template"
	"gorm.io/gorm"
)

type ForgetPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}

func (h *AuthHandler) ForgetPassword(c *gin.Context) {
	var request ForgetPasswordRequest
	if !h.ValidateRequest(c, &request) {
		return
	}

	// check if user exists
	var user model.User
	err := h.db.Where("email = ?", request.Email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			h.ResponseError(c, http.StatusBadRequest, "user not found")
		} else {
			h.ResponseError(c, http.StatusBadRequest, "error finding user: %s", err.Error())
		}
		return
	}
	// send email to user
	token := model.Token{
		Key:       uuid.New(),
		UserID:    user.ID,
		Type:      model.ResetPassword,
		ExpiresAt: time.Now().Add(time.Hour * 24),
	}

	err = h.db.Create(&token).Error
	if err != nil {
		h.ResponseError(c, http.StatusBadRequest, "error creating token: %s", err.Error())
		return
	}
	h.emailService.SendEmail(request.Email, "Forget Password", template.GetForgetPasswordTemplate(token.Key.String()))

	c.JSON(http.StatusOK, gin.H{"message": "email sent"})
}
