package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/riad/simple_auth/src/http/controllers"
)

func SetUserRouter(router *gin.RouterGroup) {
	users := router.Group("users")
	{
		users.GET("/", controllers.GetExample)
		users.POST("/", controllers.CreateExample)
		users.PUT("/:id", controllers.UpdateExample)
		users.DELETE("/:id", controllers.DeleteExample)
	}
}
