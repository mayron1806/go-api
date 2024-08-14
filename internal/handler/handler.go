package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mayron1806/go-api/config"
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
