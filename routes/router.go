package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zennon-sml/go-jwt-react/controllers"
)

func ConfigRoutes(router *gin.Engine) {
	main := router.Group("vasco/v1")
	{
		user := main.Group("user")
		{
			user.GET("/", controllers.Users)
			user.POST("/register", controllers.Register)
			user.POST("/login", controllers.Login)
		}
	}
}
