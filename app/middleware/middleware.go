package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"theAmazingNotificator/app/security"
)

func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		token, err := security.GetTokenData(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if token.Email == "" || tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{})
			c.Abort()
			return
		}

		c.Set("id", token.Id)
		c.Set("name", token.Name)
		c.Set("last_name", token.LastName)
		c.Set("email", token.Email)
	}
}
