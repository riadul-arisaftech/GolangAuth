package gapi

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/riad/simple_auth/src/db/sqlc"
	"github.com/riad/simple_auth/src/gapi/pb"
	"github.com/riad/simple_auth/src/token"
	"github.com/riad/simple_auth/src/util"
)

type Server struct {
	pb.UnimplementedSimpleAuthServer
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

	return server, nil
}
