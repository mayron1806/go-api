package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mayron1806/go-api/internal/helper"
	"github.com/mayron1806/go-api/internal/model"
	"github.com/mayron1806/go-api/internal/template"
)

type CreateUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required,gte=3,lte=50"`
	Password string `json:"password" validate:"required,gte=6,lte=50"`
}

func (h *AuthHandler) CreateUser(c *gin.Context) {
	var request CreateUserRequest
	if !h.ValidateRequest(c, &request) {
		return
	}

	hashedPassword, err := helper.HashPassword(request.Password)
	if err != nil {
		h.ResponseError(c, http.StatusBadRequest, "error hashing password: %s", err.Error())
		return
	}
	user := model.User{
		Email:     request.Email,
		Name:      request.Name,
		Password:  hashedPassword,
		Challenge: model.UserChallengeVerifyEmail,
	}
	err = h.db.Create(&user).Error
	if err != nil {
		h.ResponseError(c, http.StatusBadRequest, "error creating user: %s", err.Error())
		return
	}

	h.Logger.Infof("user created: %s", request.Email)
	token := model.Token{
		Key:       uuid.New(),
		UserID:    user.ID,
		Type:      model.ActiveAccount,
		ExpiresAt: time.Now().Add(time.Hour * 24),
	}
	err = h.db.Create(&token).Error
	if err != nil {
		h.ResponseError(c, http.StatusBadRequest, "error creating token: %s", err.Error())
		return
	}
	h.emailService.SendEmail(request.Email, "User created", template.GetActiveAccountTemplate(token.Key.String()))

	c.Status(http.StatusCreated)
}
