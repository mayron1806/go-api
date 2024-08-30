package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mayron1806/go-api/internal/services"
)

func JWTAuthMiddleware(jwtService *services.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extrai o token JWT do cabeçalho Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		// O formato esperado é "Bearer <token>"
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
			c.Abort()
			return
		}

		// Valida o token
		claims, err := jwtService.ValidateJWT(tokenString)
		if err != nil {
			println(err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: " + err.Error()})
			c.Abort()
			return
		}
		if claims.Type != services.ACCESS_TOKEN {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		// Adiciona os claims ao contexto para serem usados nos handlers
		c.Set("claims", claims)

		// Prossegue para o próximo handler
		c.Next()
	}
}
