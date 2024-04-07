package gapi

import (
	"context"
	"errors"

	pb "github.com/Streamfair/common_proto/TokenService/pb/refresh_token"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (server *Server) DeleteRefreshTokenById(ctx context.Context, req *pb.DeleteRefreshTokenByIdRequest) (*emptypb.Empty, error) {
	idParam := req.GetId()

	// Perform field validation
	if idParam <= 0 {
		violation := (&CustomError{
			StatusCode: codes.InvalidArgument,
		}).WithDetails("id", errors.New("id is required"))
		return nil, invalidArgumentError(violation)
	}

	// Verify the token exists in the database
	_, err := server.store.GetRefreshTokenById(ctx, idParam)
	if err != nil {
		// Handle database errors
		return nil, handleDatabaseError(err)
	}

	// Delete the token from the database
	err = server.store.DeleteRefreshTokenById(ctx, idParam)
	if err != nil {
		// Handle database errors
		return nil, handleDatabaseError(err)
	}

	return &emptypb.Empty{}, nil
}
