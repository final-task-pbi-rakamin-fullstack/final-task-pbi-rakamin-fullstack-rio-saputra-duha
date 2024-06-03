package helper

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckCookie(c *gin.Context, cookieName string) bool {
	if cookieName == "UserData" {
		cookie, err := c.Cookie(cookieName)
		if err != nil {
			// If the cookie is not found, return an error response
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "Unauthorized: Please login",
			})
			return false
		}
		ValidateToken(cookie)
		return true
	}
	return false
}

func SetCookie(c *gin.Context, token string) {
	// Define the cookie properties
	cookieName := "UserData"
	maxAge := 3600 // Cookie validity in seconds (1 hour)
	path := "/"
	domain := "localhost" // Change this to your domain
	secure := false       // Set to true if using HTTPS
	httpOnly := true      // Prevent JavaScript from accessing the cookie

	// Set the cookie in the response
	c.SetCookie(cookieName, token, maxAge, path, domain, secure, httpOnly)
}
