package gapi

import (
	"context"
	"errors"

	db "github.com/riad/simple_auth/src/db/sqlc"
	"github.com/riad/simple_auth/src/gapi/pb"
	"github.com/riad/simple_auth/src/http/helpers"
	"github.com/riad/simple_auth/src/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	user, err := server.Store.GetUser(ctx, req.GetEmail())
	if err != nil {
		if errors.Is(err, helpers.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to find user")
	}

	err = util.CheckPassword(req.Password, user.Password)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "incorrect password")
	}

	accessToken, accessPayload, err := server.TokenMaker.CreateToken(user.Email, server.Config.Token.AccessTokenDuration)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create access token")
	}

	refreshToken, refreshPayload, err := server.TokenMaker.CreateToken(user.Email, server.Config.Token.RefreshTokenDuration)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create refresh token")
	}

	session, err := server.Store.CreateSession(ctx, db.CreateSessionParams{
		ID:           refreshPayload.ID,
		UserID:       user.ID,
		RefreshToken: refreshToken,
		UserAgent:    "",
		ClientIp:     "",
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to find user")
	}

	rsp := &pb.LoginUserResponse{
		User:                  convertUser(user),
		SessionId:             session.ID.String(),
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  timestamppb.New(accessPayload.ExpiredAt),
		RefreshTokenExpiresAt: timestamppb.New(refreshPayload.ExpiredAt),
	}
	return rsp, nil
}
