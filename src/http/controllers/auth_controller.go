package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/riad/simple_auth/src/db/sqlc"
	"github.com/riad/simple_auth/src/http/helpers"
	"github.com/riad/simple_auth/src/http/models"
	"github.com/riad/simple_auth/src/token"
	"github.com/riad/simple_auth/src/util"
)

type Server struct {
	Config     util.Config
	Store      db.Store
	TokenMaker token.Maker
	Router     *gin.Engine
}

func CreateUser(ctx *gin.Context, store db.Store) {
	var req models.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
		return
	}
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		FullName: req.FullName,
		Email:    req.Email,
		Password: hashedPassword,
		Status:   "active",
	}

	user, err := store.CreateUser(ctx, arg)
	if err != nil {
		if helpers.ErrorCode(err) == helpers.UniqueViolation {
			ctx.JSON(http.StatusForbidden, helpers.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse(err))
		return
	}

	rsp := models.NewUserResponse(user)
	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) LoginUser(ctx *gin.Context) {
	var req models.LoginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
		return
	}

	user, err := store.GetUser(ctx, req.Email)
	if err != nil {
		if errors.Is(err, helpers.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, helpers.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse(err))
		return
	}

	err = util.CheckPassword(req.Password, user.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, helpers.ErrorResponse(err))
		return
	}

	accessToken, accessPayload, err := maker.CreateToken(user.Email)
}

// UpdateExample handles the PUT request for the example endpoint
func UpdateExample(c *gin.Context) {
	// Handle your logic here
	c.JSON(http.StatusOK, gin.H{
		"message": "Update Example",
	})
}

// DeleteExample handles the DELETE request for the example endpoint
func DeleteExample(c *gin.Context) {
	// Handle your logic here
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Example",
	})
}
