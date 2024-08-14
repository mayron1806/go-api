package helper

import "github.com/gin-gonic/gin"

// GetUserID retrieves the user ID from the Gin context claims.
//
// Parameter c is a pointer to the Gin context.
// Returns the user ID as a uint.
func GetUserID(c *gin.Context) uint {
	claims := GetClaims(c)
	if claims == nil {
		return 0
	}
	return claims.UserID
}
