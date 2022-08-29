package main

import (
	"jwt-authentication/config"
	"jwt-authentication/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialise database
	config.Connect()
	// Initialise router
	router := gin.Default()
	routes.Routes(router)

	router.Run()
}
