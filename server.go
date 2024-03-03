package main

import (
	"fmt"
	"net"
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

func getIP() {
	addr, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, addr := range addr {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			// check if IPv4 or IPv6 is not nil
			if ipnet.IP.To4() != nil || ipnet.IP.To16 != nil {
				// print available addresses
				fmt.Print("Server is running on: ")
				fmt.Println(ipnet.IP.String())
			}
		}
	}
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

	getIP()

	r.POST("/api/v1/registration", admin.Registration)
	r.POST("/api/v1/login", admin.Login)

	r.GET("/api/v1/tester", controllers.Tester)

	r.Run(fmt.Sprintf(":%s", port))
}
