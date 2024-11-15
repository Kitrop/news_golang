package middleware

import (
	"fmt"
	"net/http"

	"news-go/utils"

	"github.com/gin-gonic/gin"
)

// Autification middleware checks for a valid access token in the request cookies.
func Autification(c *gin.Context) {
	accessToken, err := c.Cookie("accessToken")
	if err != nil || accessToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Missing access token"})
		return
	}

	claims, err := utils.ValidateJWT(accessToken)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Invalid access token"})
		return
	}

	// Optionally, add user information to the context for later use.
    c.Set("user_id", claims.ID) // Example: Adding user ID to the context
    c.Set("username", claims.Username) // Example: Adding username to the context

	c.Next()
}

// UnAuthorized middleware allows access only if no access token is present.
func UnAuthorized(c *gin.Context) {
	accessToken, err := c.Cookie("accessToken")
	if err == nil && accessToken != "" {
		_, err := utils.ValidateJWT(accessToken)
		if err == nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Unauthorized: User already logged in"})
			return
		}
	}
	c.Next()
}
