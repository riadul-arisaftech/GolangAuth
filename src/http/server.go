package http

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
	db "github.com/riad/simple_auth/src/db/sqlc"
	"github.com/riad/simple_auth/src/token"
	"github.com/riad/simple_auth/src/util"
)

type Server struct {
	Config     util.Config
	Store      db.Store
	TokenMaker token.Maker
	Router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.Token.SymetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		Config:     config,
		Store:      store,
		TokenMaker: tokenMaker,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("status", ValidStatus)
	}

	server.setupRouter()

	return server, nil
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.Router.Run(address)
}

var ValidStatus validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if status, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsSupportedStatus(status)
	}
	return false
}
