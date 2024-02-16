package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var input User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(input)
	if input.Username != "idam_kusdiana" {
		if input.Password != "super_sapar" {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid password and username"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid username"})
		return
	}
	token, err := generateToken(input.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":   "success loging in",
		"token": token,
	})
}
