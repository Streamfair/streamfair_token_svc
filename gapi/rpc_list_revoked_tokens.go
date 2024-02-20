package gapi

import (
	"context"

	db "github.com/Streamfair/streamfair_token_svc/db/sqlc"
	pb "github.com/Streamfair/streamfair_token_svc/pb/token"
	"github.com/Streamfair/streamfair_token_svc/validator"
	"google.golang.org/grpc/codes"
)

func (server *Server) ListRevokedTokens(ctx context.Context, req *pb.ListRevokedTokensRequest) (*pb.ListRevokedTokensResponse, error) {
	// Validate the request
	violations := validateListRevokedTokensRequest(req)
	if len(violations) > 0 {
		return nil, invalidArgumentErrors(violations)
	}

	tokens, err := server.store.ListRevokedTokens(ctx, db.ListRevokedTokensParams{
		Limit:  req.GetLimit(),
		Offset: req.GetOffset(),
	})
	if err != nil {
		// Handle database errors
		return nil, handleDatabaseError(err)
	}

	rsp := &pb.ListRevokedTokensResponse{
		Tokens: convertTokenList(tokens),
	}

	return rsp, nil
}

func validateListRevokedTokensRequest(req *pb.ListRevokedTokensRequest) (violations []*CustomError) {
	if err := validator.ValidateLimit(req.GetLimit()); err != nil {
		violations = append(violations, (&CustomError{
			StatusCode: codes.OutOfRange,
		}).WithDetails("limit", err))
	}

	if err := validator.ValidateOffset(req.GetOffset()); err != nil {
		violations = append(violations, (&CustomError{
			StatusCode: codes.OutOfRange,
		}).WithDetails("offset", err))
	}

	return violations
}
