package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zennon-sml/go-jwt-react/database"
	"github.com/zennon-sml/go-jwt-react/routes"
)

func main() {
	database.MKCon()

	server := gin.Default()

	routes.ConfigRoutes(server)
	server.Run(":8080")
}
