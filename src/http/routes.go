package http

import (
	"github.com/gin-gonic/gin"
	"github.com/riad/simple_auth/src/http/controllers"
	"github.com/riad/simple_auth/src/http/middlewares"
)

func (server *Server) setupRouter() {
	router := gin.Default()

	authController := controllers.AuthController{Store: server.Store, Maker: server.TokenMaker, Config: server.Config}

	api := router.Group("api")

	auth := api.Group("auth")
	{
		auth.POST("/signup", authController.CreateUser)
		auth.POST("/signin", authController.LoginUser)
	}

	users := api.Group("users").Use(middlewares.AuthMiddleware(server.TokenMaker))
	{
		users.PUT("/:id", controllers.UpdateExample)
		users.DELETE("/:id", controllers.DeleteExample)
	}

	server.Router = router
}
