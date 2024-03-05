package gapi

import (
	"context"
	"errors"

	pb "github.com/Streamfair/streamfair_token_svc/pb/refresh_token"
	"google.golang.org/grpc/codes"
)

func (server *Server) GetRefreshTokenByValue(ctx context.Context, req *pb.GetRefreshTokenByValueRequest) (*pb.GetRefreshTokenByValueResponse, error) {
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

	rsp := &pb.GetRefreshTokenByValueResponse{
		RefreshToken: convertRefreshToken(refreshToken),
	}

	return rsp, nil
}
