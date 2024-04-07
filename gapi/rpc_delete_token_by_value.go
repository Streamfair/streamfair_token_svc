package gapi

import (
	"context"
	"errors"

	pb "github.com/Streamfair/common_proto/TokenService/pb/token"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (server *Server) DeleteTokenByValue(ctx context.Context, req *pb.DeleteTokenByValueRequest) (*emptypb.Empty, error) {
	tokenParam := req.GetToken()

	// Perform field validation
	if tokenParam <= "" {
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

	// Delete the token from the database
	err = server.store.DeleteTokenByValue(ctx, token.Token)
	if err != nil {
		// Handle database errors
		return nil, handleDatabaseError(err)
	}

	return &emptypb.Empty{}, nil
}
