package routes

import (
	"github.com/gin-gonic/gin"
	db "github.com/riad/simple_auth/src/db/sqlc"
	"github.com/riad/simple_auth/src/http/controllers"
	"github.com/riad/simple_auth/src/http/middlewares"
	"github.com/riad/simple_auth/src/token"
	"github.com/riad/simple_auth/src/util"
)

type Server struct {
	Config     util.Config
	Store      db.Store
	TokenMaker token.Maker
	Router     *gin.Engine
}

func (server *Server) SetAuthRouter(router *gin.RouterGroup) {
	auth := router.Group("auth")
	{
		auth.POST("/signup", controllers.CreateUser)
		auth.POST("/signin", func(ctx *gin.Context) {
			controllers.LoginUser(ctx, store, maker)
		})
	}

	users := router.Group("users").Use(middlewares.AuthMiddleware(maker))
	{
		users.PUT("/:id", controllers.UpdateExample)
		users.DELETE("/:id", controllers.DeleteExample)
	}
}
