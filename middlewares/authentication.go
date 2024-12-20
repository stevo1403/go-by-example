package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/stevo1403/go-by-example/apps/user"
)

func UserAuthMiddleware(c *gin.Context) {
	// Extract Authorization header
	authorizationHeader := c.Request.Header.Get("Authorization")

	// Check if Authorization header is empty
	if authorizationHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
		return
	}

	// Extract Bearer token from header value
	bearerToken := strings.TrimPrefix(authorizationHeader, "Bearer ")

	// Check if token is empty
	if bearerToken == ""{
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
		return
	}

	// Check if token is valid
	isTokenValid, err := user.User{}.VerifyToken(bearerToken)
	if isTokenValid && err == nil {
		c.Next()
		return
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
}
