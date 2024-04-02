package gapi

import (
	"context"
	"errors"

	pb "github.com/Streamfair/streamfair_token_svc/common_proto/TokenService/pb/token"
	"google.golang.org/grpc/codes"
)

func (server *Server) GetTokenById(ctx context.Context, req *pb.GetTokenByIdRequest) (*pb.GetTokenByIdResponse, error) {
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

	rsp := &pb.GetTokenByIdResponse{
		Token: convertToken(token),
	}

	return rsp, nil
}
