package gapi

import (
	"context"
	"errors"

	pb "github.com/Streamfair/streamfair_token_svc/pb/token"
	"google.golang.org/grpc/codes"
)

func (server *Server) GetTokenByValue(ctx context.Context, req *pb.GetTokenByValueRequest) (*pb.GetTokenByValueResponse, error) {
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

	rsp := &pb.GetTokenByValueResponse{
		Token: convertToken(token),
	}

	return rsp, nil
}