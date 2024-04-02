package gapi

import (
	"context"
	"errors"

	pb "github.com/Streamfair/streamfair_token_svc/common_proto/TokenService/pb/refresh_token"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (server *Server) RevokeRefreshTokenByValue(ctx context.Context, req *pb.RevokeRefreshTokenByValueRequest) (*emptypb.Empty, error) {
	tokenParam := req.GetRefreshToken()

	// Perform field validation
	if tokenParam == "" {
		violation := (&CustomError{
			StatusCode: codes.InvalidArgument,
		}).WithDetails("token", errors.New("token is required"))
		return nil, invalidArgumentError(violation)
	}

	// Verify the token exists in the database
	refreshToken, err := server.store.GetRefreshTokenByValue(ctx, tokenParam)
	if err != nil {
		// Handle database errors
		return nil, handleDatabaseError(err)
	}

	if refreshToken.Revoked {
		// Token is already revoked
		violation := (&CustomError{
			StatusCode: codes.InvalidArgument,
		}).WithDetails("token", errors.New("token is already revoked"))
		return nil, invalidArgumentError(violation)
	}

	// Revoke the token in the database
	err = server.store.RevokeRefreshTokenByValue(ctx, refreshToken.Token)
	if err != nil {
		// Handle database errors
		return nil, handleDatabaseError(err)
	}

	return &emptypb.Empty{}, nil
}