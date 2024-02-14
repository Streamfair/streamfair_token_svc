package gapi

import (
	db "github.com/Streamfair/streamfair_token_svc/db/sqlc"
	pb "github.com/Streamfair/streamfair_token_svc/pb/token"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertToken(token db.TokenSvcToken) *pb.Token {
	return &pb.Token{
		Id:        token.ID,
		UserId:    token.UserID,
		TokenType: token.TokenType.String,
		Token:     token.Token,
		ExpiresAt: timestamppb.New(token.ExpiresAt),
		CreatedAt: timestamppb.New(token.CreatedAt),
	}
}
