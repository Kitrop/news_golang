package middleware

import (
	"net/http"

	"news-go/utils"

	"github.com/gin-gonic/gin"
)

// Middleware for authentication, checks if the accessToken exists and is valid 
func Autification(c *gin.Context) {
	accessToken, err := c.Cookie("accessToken")
	if err != nil || accessToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Missing access token"})
		return
	}

	claims, err := utils.ValidateJWT(accessToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Invalid access token"})
		return
	}
	
	if claims.Role == "BANNED" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: user banned"})
		return
	}

  c.Set("user_id", claims.ID)
  c.Set("username", claims.Username)

	c.Next()
}

// UnAuthorized middleware allows access only if no access token is present
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
