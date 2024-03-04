package controllers

import (
	"net/http"

	jwt_auth "github.com/faizallmaullana/test-koyeb/Authentication"
	"github.com/gin-gonic/gin"
)

func Tester(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	profileData, err := jwt_auth.JWTClaims(tokenString, "guru")
	if err == nil {
		if profileData["status"] != "Authorized" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "You are unauthorized to access this feature"})
			return
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Error parsing token authorization"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": profileData})
}
