package gapi

import (
	db "github.com/Streamfair/streamfair_token_svc/db/sqlc"
	pb_rtoken "github.com/Streamfair/common_proto/TokenService/pb/refresh_token"
	pb_token "github.com/Streamfair/common_proto/TokenService/pb/token"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertToken(token db.TokenSvcToken) *pb_token.Token {
	return &pb_token.Token{
		Id:        token.ID,
		UserId:    token.UserID,
		Token:     token.Token,
		Revoked:   token.Revoked,
		ExpiresAt: timestamppb.New(token.ExpiresAt),
		CreatedAt: timestamppb.New(token.CreatedAt),
		UpdatedAt: timestamppb.New(token.UpdatedAt),
	}
}

func convertTokenList(tokens []db.TokenSvcToken) []*pb_token.Token {
	var tokenList []*pb_token.Token
	for _, token := range tokens {
		tokenList = append(tokenList, &pb_token.Token{
			Id:        token.ID,
			UserId:    token.UserID,
			Token:     token.Token,
			Revoked:   token.Revoked,
			ExpiresAt: timestamppb.New(token.ExpiresAt),
			CreatedAt: timestamppb.New(token.CreatedAt),
			UpdatedAt: timestamppb.New(token.UpdatedAt),
		})
	}
	return tokenList
}

func convertRefreshToken(token db.TokenSvcRefreshToken) *pb_rtoken.RefreshToken {
	return &pb_rtoken.RefreshToken{
		Id:        token.ID,
		UserId:    token.UserID,
		Token:     token.Token,
		Revoked:   token.Revoked,
		ExpiresAt: timestamppb.New(token.ExpiresAt),
		CreatedAt: timestamppb.New(token.CreatedAt),
		UpdatedAt: timestamppb.New(token.UpdatedAt),
	}
}

func convertRefreshTokenList(refreshTokens []db.TokenSvcRefreshToken) []*pb_rtoken.RefreshToken {
	var refreshTokenList []*pb_rtoken.RefreshToken
	for _, refreshToken := range refreshTokens {
		refreshTokenList = append(refreshTokenList, &pb_rtoken.RefreshToken{
			Id:        refreshToken.ID,
			UserId:    refreshToken.UserID,
			Token:     refreshToken.Token,
			Revoked:   refreshToken.Revoked,
			ExpiresAt: timestamppb.New(refreshToken.ExpiresAt),
			CreatedAt: timestamppb.New(refreshToken.CreatedAt),
			UpdatedAt: timestamppb.New(refreshToken.UpdatedAt),
		})
	}
	return refreshTokenList
}