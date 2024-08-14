package auth

import (
	"github.com/mayron1806/go-api/config"
	"github.com/mayron1806/go-api/internal/handler"
	"github.com/mayron1806/go-api/internal/services"
	"gorm.io/gorm"
)

type AuthHandler struct {
	*handler.Handler
	emailService *services.EmailService
	jwtService   *services.JWTService
	db           *gorm.DB
}

func NewAuthHandler() (*AuthHandler, error) {
	logger := config.GetLogger("Auth Handler")
	emailService := services.NewEmailService()
	jwtService := services.NewJWTService()
	db := config.GetDatabase()

	handler := &AuthHandler{
		Handler:      handler.NewHandler(logger),
		emailService: emailService,
		jwtService:   jwtService,
		db:           db,
	}
	return handler, nil
}
