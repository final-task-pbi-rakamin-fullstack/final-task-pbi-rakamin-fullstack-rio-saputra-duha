package middleware

import (
	"github.com/gin-gonic/gin"
	"goAPI/helper"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//midleware
		tokenString, _ := c.Cookie("UserData")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Message": "Unauthorized",
			})
			c.Abort()
			return
		}

		userUUID, err := helper.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Message": err.Error(),
			})
			c.Abort()
			return
		}

		c.Set("userUuid", *userUUID)
		c.Next()
	}
}
