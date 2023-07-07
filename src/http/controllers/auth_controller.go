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

type AuthController struct {
	Store  db.Store
	Maker  token.Maker
	Config util.Config
}

func (ctrl AuthController) CreateUser(ctx *gin.Context) {
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

	user, err := ctrl.Store.CreateUser(ctx, arg)
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

func (ctrl AuthController) LoginUser(ctx *gin.Context) {
	var req models.LoginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
		return
	}

	user, err := ctrl.Store.GetUser(ctx, req.Email)
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

	accessToken, accessPayload, err := ctrl.Maker.CreateToken(user.Email, ctrl.Config.Token.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse(err))
		return
	}

	refreshToken, refreshPayload, err := ctrl.Maker.CreateToken(user.Email, ctrl.Config.Token.RefreshTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse(err))
		return
	}

	session, err := ctrl.Store.CreateSession(ctx, db.CreateSessionParams{
		ID:           refreshPayload.ID,
		UserID:       user.ID,
		RefreshToken: refreshToken,
		UserAgent:    ctx.Request.UserAgent(),
		ClientIp:     ctx.ClientIP(),
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse(err))
		return
	}

	rsp := models.LoginUserResponse{
		SessionID:             session.ID,
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
		User:                  models.NewUserResponse(user),
	}

	ctx.JSON(http.StatusOK, rsp)
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
