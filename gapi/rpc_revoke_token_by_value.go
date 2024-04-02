package gapi

import (
	"context"
	"errors"

	pb "github.com/Streamfair/streamfair_token_svc/common_proto/TokenService/pb/token"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (server *Server) RevokeTokenByValue(ctx context.Context, req *pb.RevokeTokenByValueRequest) (*emptypb.Empty, error) {
	tokenParam := req.GetToken()

	// Perform field validation
	if tokenParam == "" {
		violation := (&CustomError{
			StatusCode: codes.InvalidArgument,
		}).WithDetails("token", errors.New("token is required"))
		return nil, invalidArgumentError(violation)
	}

	// Verify the token exists in the database
	token, err := server.store.GetTokenByValue(ctx, tokenParam)
	if err != nil {
		// Handle database errors
		return nil, handleDatabaseError(err)
	}

	if token.Revoked {
		// Token is already revoked
		violation := (&CustomError{
			StatusCode: codes.InvalidArgument,
		}).WithDetails("token", errors.New("token is already revoked"))
		return nil, invalidArgumentError(violation)
	}

	// Revoke the token in the database
	err = server.store.RevokeTokenByValue(ctx, token.Token)
	if err != nil {
		// Handle database errors
		return nil, handleDatabaseError(err)
	}

	return &emptypb.Empty{}, nil
}
