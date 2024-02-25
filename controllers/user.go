package controllers

import (
	"net/http"

	"github.com/faizallmaullana/test-koyeb/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUser(c *gin.Context) {
	var User models.Users
	if err := models.DB.First(&User).Error; err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": User})
}

func CreateUser(c *gin.Context) {
	id := uuid.New().String()

	data := models.Users{
		ID:       id,
		Username: "faizal",
	}

	models.DB.Create(data)
	c.JSON(http.StatusCreated, gin.H{"data": data})
}
