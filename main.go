package main

import (
	"github.com/BerkAkipek/rest_api_go/db"
	"github.com/BerkAkipek/rest_api_go/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
