package gapi

import (
	db "github.com/riad/simple_auth/src/db/sqlc"
	"github.com/riad/simple_auth/src/gapi/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.User) *pb.User {
	return &pb.User{
		FullName:          user.FullName,
		Email:             user.Email,
		Status:            user.Status,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt:         timestamppb.New(user.CreatedAt),
	}
}
