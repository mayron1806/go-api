package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mayron1806/go-api/config"
	"github.com/mayron1806/go-api/internal/services"
)

type cookiesConfig struct {
	domain   string
	path     string
	secure   bool
	httpOnly bool
}
type Handler struct {
	Logger        *config.Logger
	cookiesConfig *cookiesConfig
}

func NewHandler(logger *config.Logger) *Handler {
	env := config.GetEnv()
	handler := &Handler{
		Logger: logger,
		cookiesConfig: &cookiesConfig{
			domain:   env.COOKIES_DOMAIN,
			path:     env.COOKIES_PATH,
			httpOnly: env.COOKIES_HTTP_ONLY,
			secure:   env.COOKIES_SECURE,
		},
	}
	return handler
}
func (h Handler) SetCookie(ctx *gin.Context, name, value string, maxAge int) {
	ctx.SetCookie(
		name,
		value,
		maxAge,
		h.cookiesConfig.path,
		h.cookiesConfig.domain,
		h.cookiesConfig.secure,
		h.cookiesConfig.httpOnly,
	)
}
func (h Handler) ValidateRequest(ctx *gin.Context, request interface{}) bool {
	if err := ctx.BindJSON(request); err != nil {
		h.Logger.Errorf("error validating request: %s", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return false
	}
	validator := validator.New()
	err := validator.Struct(request)
	if err != nil {
		h.Logger.Errorf("error validating request body: %s", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return false
	}
	return true
}
func (h Handler) ResponseError(ctx *gin.Context, status int, format string, a ...any) {
	var errorMessage = fmt.Sprintf(format, a...)
	h.Logger.Errorf("error: %s", errorMessage)
	ctx.JSON(status, gin.H{"error": errorMessage})
}
func (h Handler) GetClaims(c *gin.Context) *services.JWTClaims {
	contextClaims, exists := c.Get("claims")
	if !exists {
		return nil
	}

	claims, ok := contextClaims.(services.JWTClaims)
	if !ok {
		return nil
	}
	return &claims
}
func (h Handler) GetUserID(c *gin.Context) uint {
	claims := h.GetClaims(c)
	if claims == nil {
		return 0
	}
	return claims.UserID
}
