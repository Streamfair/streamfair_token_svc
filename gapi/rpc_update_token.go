package gapi

import (
	"context"
	"errors"

	db "github.com/Streamfair/streamfair_token_svc/db/sqlc"
	pb "github.com/Streamfair/streamfair_token_svc/pb/token"
	"google.golang.org/grpc/codes"
)

func (server *Server) UpdateToken(ctx context.Context, req *pb.UpdateTokenRequest) (*pb.UpdateTokenResponse, error) {
	// Validate the request
	if req.GetUserId() <= 0 {
		violation := (&CustomError{
			StatusCode: codes.InvalidArgument,
		}).WithDetails("id", errors.New("id is required"))
		return nil, invalidArgumentError(violation)
	}

	// Update the token in the database
	dbToken, err := server.store.UpdateToken(ctx, db.UpdateTokenParams{
		UserID: req.GetId(),
		ID:     req.GetId(),
	})
	if err != nil {
		return nil, handleDatabaseError(err)
	}

	rsp := &pb.UpdateTokenResponse{
		Token: convertToken(dbToken),
	}

	return rsp, nil
}
