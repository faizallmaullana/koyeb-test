package main

import (
	"fmt"
	"os"

	"github.com/faizallmaullana/test-koyeb/controllers"
	"github.com/faizallmaullana/test-koyeb/controllers/admin"
	"github.com/faizallmaullana/test-koyeb/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var corsConfig = cors.DefaultConfig()

func init() {
	// allow all origins
	corsConfig.AllowAllOrigins = true
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	// Use cors middleware
	models.ConnectToDatabase()
	r.Use(cors.New(corsConfig))
	fmt.Println("server is running")

	r.POST("/api/v1/registration", admin.Registration)
	r.POST("/api/v1/login", admin.Login)

	r.GET("/api/v1/tester", controllers.Tester)

	r.Run(fmt.Sprintf(":%s", port))
}
