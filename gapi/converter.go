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
		Revoked:   token.Revoked,
		ExpiresAt: timestamppb.New(token.ExpiresAt),
		CreatedAt: timestamppb.New(token.CreatedAt),
		UpdatedAt:  timestamppb.New(token.UpdatedAt),
	}
}

func convertTokenList(tokens []db.TokenSvcToken) []*pb.Token {
	var tokenList []*pb.Token
	for _, token := range tokens {
		tokenList = append(tokenList, &pb.Token{
			Id:        token.ID,
			UserId:    token.UserID,
			TokenType: token.TokenType.String,
			Token:     token.Token,
			Revoked:   token.Revoked,
			ExpiresAt: timestamppb.New(token.ExpiresAt),
			CreatedAt: timestamppb.New(token.CreatedAt),
			UpdatedAt:  timestamppb.New(token.UpdatedAt),
		})
	}
	return tokenList
}
