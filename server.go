package main

import (
	"fmt"
	"os"

	"github.com/faizallmaullana/test-koyeb/controller"
	"github.com/faizallmaullana/test-koyeb/jsonData"
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
	r.Use(cors.New(corsConfig))

	models.ConnectToDatabase()

	fmt.Println("server is running")

	r.GET("api/v1/kopisore/tables", models.ReadTable)
	r.DELETE("api/v1/kopisore/tables", models.DropTable)

	r.GET("api/v1/kopisore/menus", controller.GetAllMenu)
	r.POST("api/v1/kopisore/menus", jsonData.ReadJson)

	r.POST("api/v1/kopisore/login", controller.Login)

	r.Run(fmt.Sprintf(":%s", port))
}
