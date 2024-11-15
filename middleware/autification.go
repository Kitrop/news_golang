package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"news-go/utils"
)

func Autification(c *gin.Context) {
	accessToken, err := c.Cookie("accessToken")

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authorized"})
		return
	}

	_, err = utils.ValidateJWT(accessToken)

	if err != nil {
		c.SetCookie("accessToken", "", -1, "/", "localhost", false, true)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authorized"})
		return
	}

	c.Next()
}

func UnAuthorized(c *gin.Context) {
	accessToken, err := c.Cookie("accessToken")

	if err != nil {
		c.SetCookie("accessToken", "", -1, "/", "localhost", false, true)
		c.Next()
		return
	}

	_, err = utils.ValidateJWT(accessToken)

	if err != nil {
		c.SetCookie("accessToken", "", -1, "/", "localhost", false, true)
		c.Next()
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": "user already authorized"})
}
