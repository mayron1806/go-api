package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mayron1806/go-api/config"
	"github.com/mayron1806/go-api/internal/services"
)

type Handler struct {
	Logger *config.Logger
}

func NewHandler(logger *config.Logger) *Handler {
	handler := &Handler{
		Logger: logger,
	}
	return handler
}
func (h *Handler) ValidateRequest(ctx *gin.Context, request interface{}) bool {
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
func (h *Handler) ResponseError(ctx *gin.Context, status int, format string, a ...any) {
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
