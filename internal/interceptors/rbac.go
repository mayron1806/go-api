package interceptors

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mayron1806/go-api/config"
	"github.com/mayron1806/go-api/internal/model"
	"github.com/mayron1806/go-api/internal/services"
)

func RBAC(next gin.HandlerFunc, requiredPermissions ...model.Permission) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extrai o token JWT do cabeçalho Authorization
		var claims services.JWTClaims
		if claimsAny, exists := c.Get("claims"); !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		} else {
			// Faça a type assertion para garantir que o valor é do tipo services.JWTClaims
			var ok bool
			if claims, ok = claimsAny.(services.JWTClaims); !ok {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
				c.Abort()
				return
			}
		}
		formattedRequiredPermissions := make([]model.Permission, len(requiredPermissions))

		if organizationId, ok := c.Params.Get("organizationId"); ok {
			for i, permission := range requiredPermissions {
				formattedRequiredPermissions[i] = model.Permission(strings.Replace(permission.String(), "{organizationId}", organizationId, -1))
			}
		}
		logger := config.GetLogger("RBAC")
		logger.Debugf("requiredPermissions: %v", formattedRequiredPermissions)
		logger.Debugf("user permissions: %v", claims.Permissions)
		allPermissionsOk := true
		for _, permission := range formattedRequiredPermissions {
			permissionOk := false
			for _, rolePermission := range claims.Permissions {
				if rolePermission == permission {
					permissionOk = true
					break
				}
			}
			if !permissionOk {
				allPermissionsOk = false
				break
			}
		}

		if !allPermissionsOk {
			c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
			c.Abort()
			return
		}

		next(c)
	}
}
