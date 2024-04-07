package gapi

import (
	"context"
	"errors"
	"fmt"

	pb "github.com/Streamfair/common_proto/TokenService/pb/refresh_token"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (server *Server) RevokeRefreshTokenById(ctx context.Context, req *pb.RevokeRefreshTokenByIdRequest) (*emptypb.Empty, error) {
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

	if refreshToken.Revoked {
		// Token is already revoked
		violation := (&CustomError{
			StatusCode: codes.InvalidArgument,
		}).WithDetails("id", fmt.Errorf("token with id '%d' is already revoked", idParam))
		return nil, invalidArgumentError(violation)
	}

	// Revoke the token in the database
	err = server.store.RevokeRefreshTokenById(ctx, idParam)
	if err != nil {
		// Handle database errors
		return nil, handleDatabaseError(err)
	}

	return &emptypb.Empty{}, nil
}
