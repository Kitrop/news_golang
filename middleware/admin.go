package middleware

import (
	"net/http"
	"news-go/utils"

	"github.com/gin-gonic/gin"
)

// middleware для проверки, того что пользователь является админом
func СheckIsAdmin(c *gin.Context) {
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

	if claims.Role != "ADMIN" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "user is not an admin"})
		return
	}

	c.Next()
}