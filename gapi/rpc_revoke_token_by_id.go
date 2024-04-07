package gapi

import (
	"context"
	"errors"
	"fmt"

	pb "github.com/Streamfair/common_proto/TokenService/pb/token"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (server *Server) RevokeTokenById(ctx context.Context, req *pb.RevokeTokenByIdRequest) (*emptypb.Empty, error) {
	idParam := req.GetId()

	// Perform field validation
	if idParam <= 0 {
		violation := (&CustomError{
			StatusCode: codes.InvalidArgument,
		}).WithDetails("id", errors.New("id is required"))
		return nil, invalidArgumentError(violation)
	}

	// Verify the token exists in the database
	token, err := server.store.GetTokenById(ctx, idParam)
	if err != nil {
		// Handle database errors
		return nil, handleDatabaseError(err)
	}

	if token.Revoked {
		// Token is already revoked
		violation := (&CustomError{
			StatusCode: codes.InvalidArgument,
		}).WithDetails("id", fmt.Errorf("token with id '%d' is already revoked", idParam))
		return nil, invalidArgumentError(violation)
	}

	// Revoke the token in the database
	err = server.store.RevokeTokenById(ctx, idParam)
	if err != nil {
		// Handle database errors
		return nil, handleDatabaseError(err)
	}

	return &emptypb.Empty{}, nil
}
