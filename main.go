package main

import (
	"github.com/carlosgrillet/go-api/db"
	"github.com/carlosgrillet/go-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
