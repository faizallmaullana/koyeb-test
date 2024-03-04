package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/faizallmaullana/test-koyeb/controllers"
	"github.com/faizallmaullana/test-koyeb/controllers/admin"
	"github.com/faizallmaullana/test-koyeb/controllers/siswa"
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

	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	if port == "" {
		port = "8000"
	}

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	// Use cors middleware
	models.ConnectToDatabase()
	r.Use(cors.New(corsConfig))

	getIP()
	fmt.Printf("Port: %s \n", port)

	// auth
	r.POST("/api/v1/registration", admin.Registration)
	r.POST("/api/v1/login", admin.Login)

	// siswa
	r.POST("/api/v1/siswa", siswa.PostSiswa)
	r.POST("/api/v1/siswa/image/:id", siswa.PostImageSiswa)

	// stats
	r.POST("/api/v1/tester/:id", controllers.Tester)
	r.POST("/api/v1/tester2", controllers.Tester2)
	r.GET("/api/v1/tester", controllers.GetTester)
	r.GET("/api/v1/memory", GetMemory)

	r.GET("/api/v1/DownloadDB", admin.DownloadDB)

	r.Run(fmt.Sprintf(":%s", port))
}
