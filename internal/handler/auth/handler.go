package auth

import (
	"github.com/mayron1806/go-api/config"
	"github.com/mayron1806/go-api/internal/goauth2"
	"github.com/mayron1806/go-api/internal/goauth2/google"
	"github.com/mayron1806/go-api/internal/handler"
	"github.com/mayron1806/go-api/internal/query"
	"github.com/mayron1806/go-api/internal/services"
	"gorm.io/gorm"
)

type AuthHandler struct {
	*handler.Handler
	emailService *services.EmailService
	jwtService   *services.JWTService
	db           *gorm.DB
	queryUser    *query.QueryUser
}

func NewAuthHandler() (*AuthHandler, error) {
	env := config.GetEnv()
	if env.GOOGLE_OAUTH_ENABLED {
		goauth2.AddProvider("google", google.New(env.GOOGLE_OAUTH_CLIENT_ID, env.GOOGLE_OAUTH_CLIENT_SECRET, env.GOOGLE_OAUTH_REDIRECT_URL))
	}

	logger := config.GetLogger("Auth Handler")
	emailService := services.NewEmailService()
	jwtService := services.NewJWTService()
	db := config.GetDatabase()
	queryUser := query.NewQueryUser()

	handler := &AuthHandler{
		Handler:      handler.NewHandler(logger),
		emailService: emailService,
		jwtService:   jwtService,
		db:           db,
		queryUser:    queryUser,
	}
	return handler, nil
}
