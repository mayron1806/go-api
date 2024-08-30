package auth

import (
	"time"

	"github.com/gin-gonic/gin"
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
	authService  *services.AuthService
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
	jwtService := services.NewAuthService()
	db := config.GetDatabase()
	queryUser := query.NewQueryUser()

	handler := &AuthHandler{
		Handler:      handler.NewHandler(logger),
		emailService: emailService,
		authService:  jwtService,
		db:           db,
		queryUser:    queryUser,
	}
	return handler, nil
}
func (h *AuthHandler) SetTokenCookies(c *gin.Context, tokens services.GenerateTokensResponse) {
	h.SetCookie(c, "access-token", tokens.AccessToken.Token, int(tokens.AccessToken.ExpiresAt.Sub(time.Now()).Seconds()))
	h.SetCookie(c, "expires-at", tokens.AccessToken.ExpiresAt.Format(time.RFC3339), int(tokens.AccessToken.ExpiresAt.Sub(time.Now()).Seconds()))
	h.SetCookie(c, "refresh-token", tokens.RefreshToken.Token, int(tokens.RefreshToken.ExpiresAt.Sub(time.Now()).Seconds()))
}
