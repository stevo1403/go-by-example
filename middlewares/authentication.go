package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/stevo1403/go-by-example/apps/user"
)

func UserAuthMiddleware(c *gin.Context) {
	// Extract Authorization header
	authorizationHeader := c.Request.Header.Get("Authorization")

	// Extract Bearer token from header value
	bearerToken := strings.TrimPrefix(authorizationHeader, "Bearer ")

	fmt.Println(bearerToken)

	// Check if token is valid
	isTokenValid, err := user.User{}.VerifyToken(bearerToken)
	fmt.Println(isTokenValid, err)

	if isTokenValid && err == nil {
		c.Next()
		return
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}
}
