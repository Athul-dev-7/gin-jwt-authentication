package routes

import (
	"jwt-authentication/controllers"
	"jwt-authentication/middlewares"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/user/register", controllers.RegisterUser)
		api.POST("/token", controllers.GenerateToken)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}
}
