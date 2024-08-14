package helper

import (
	"github.com/gin-gonic/gin"
	"github.com/mayron1806/go-api/internal/services"
)

// GetClaims retrieves the JWT claims from the Gin context.
func GetClaims(c *gin.Context) *services.JWTClaims {
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
