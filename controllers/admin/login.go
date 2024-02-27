package admin

import (
	"net/http"

	jwt_auth "github.com/faizallmaullana/test-koyeb/Authentication"
	"github.com/faizallmaullana/test-koyeb/controllers"
	"github.com/faizallmaullana/test-koyeb/models"
	"github.com/gin-gonic/gin"
)

type inputLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var input inputLogin
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// username and password validation
	var user models.Authentication
	if err := models.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username not valid"})
		return
	}

	matchPassword := controllers.CheckPasswordHash(input.Password, user.Password)
	if !matchPassword {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Wrong Password"})
		return
	}

	tokenJWT, err := jwt_auth.GenerateToken(user.Username, user.ID, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       "success",
		"authorization": tokenJWT,
		"role":          user.Role,
	})
}

// input
// validate the username and password
// generate jwt token to authorization
// return data
