package main

import (
	"fmt"
	"os"

	"github.com/faizallmaullana/test-koyeb/controller"
	"github.com/faizallmaullana/test-koyeb/jsonData"
	"github.com/faizallmaullana/test-koyeb/models"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // Import PostgreSQL driver
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	models.ConnectToDatabase()

	fmt.Println("server is running")

	r.GET("api/v1/kopisore/tables", models.ReadTable)
	r.DELETE("api/v1/kopisore/tables", models.DropTable)

	r.GET("api/v1/kopisore/menus", controller.GetAllMenu)
	r.POST("api/v1/kopisore/menus", jsonData.ReadJson)

	r.Run(fmt.Sprintf(":%s", port))
}
