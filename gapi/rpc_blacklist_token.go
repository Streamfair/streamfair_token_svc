package gapi

import (
	"context"
	"errors"

	pb "github.com/Streamfair/streamfair_token_svc/pb/token"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (server *Server) BlacklistToken(ctx context.Context, req *pb.BlacklistTokenRequest) (*emptypb.Empty, error) {
	// Extract the token from the request
	tokenParam := req.GetToken()

	// Perform field validation
	if tokenParam == "" {
		fieldViolation := fieldViolation("token", errors.New("token is required"))
		violations := []*errdetails.BadRequest_FieldViolation{fieldViolation}
		return nil, invalidArgumentError(violations)
	}

	_, err := server.localTokenMaker.VerifyLocalToken(tokenParam)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token: %v", err)
	}

	// Blacklist the token in the database
	err = server.store.BlacklistToken(ctx, tokenParam)
	if err != nil {
		// Handle database errors
		return nil, handleDatabaseError(err)
	}

	// Return a success message along with the empty response
	return &emptypb.Empty{}, nil
}
