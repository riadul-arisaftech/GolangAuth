package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/riad/simple_auth/src/http/controllers"
	"github.com/riad/simple_auth/src/http/middlewares"
	"github.com/riad/simple_auth/src/token"
)

func SetUserRouter(router *gin.RouterGroup, maker token.Maker) {
	users := router.Group("users")
	{
		users.GET("/", controllers.GetExample)
		users.POST("/", controllers.CreateExample)
	}

	authRoute := users.Use(middlewares.AuthMiddleware(maker))
	{
		authRoute.PUT("/:id", controllers.UpdateExample)
		authRoute.DELETE("/:id", controllers.DeleteExample)
	}
}
