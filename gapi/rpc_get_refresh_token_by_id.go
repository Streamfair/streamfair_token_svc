package gapi

import (
	"context"
	"errors"

	pb "github.com/Streamfair/streamfair_token_svc/pb/refresh_token"
	"google.golang.org/grpc/codes"
)

func (server *Server) GetRefreshTokenById(ctx context.Context, req *pb.GetRefreshTokenByIdRequest) (*pb.GetRefreshTokenByIdResponse, error) {
	idParam := req.GetId()

	// Perform field validation
	if idParam <= 0 {
		violation := (&CustomError{
			StatusCode: codes.InvalidArgument,
		}).WithDetails("id", errors.New("id is required"))
		return nil, invalidArgumentError(violation)
	}

	// Verify the token exists in the database
	refreshToken, err := server.store.GetRefreshTokenById(ctx, idParam)
	if err != nil {
		// Handle database errors
		return nil, handleDatabaseError(err)
	}

	rsp := &pb.GetRefreshTokenByIdResponse{
		RefreshToken: convertRefreshToken(refreshToken),
	}

	return rsp, nil
}
