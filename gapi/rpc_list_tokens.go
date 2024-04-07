package gapi

import (
	"context"

	db "github.com/Streamfair/streamfair_token_svc/db/sqlc"
	pb "github.com/Streamfair/common_proto/TokenService/pb/token"
	"github.com/Streamfair/streamfair_token_svc/validator"
	"google.golang.org/grpc/codes"
)

func (server *Server) ListTokens(ctx context.Context, req *pb.ListTokensRequest) (*pb.ListTokensResponse, error) {
	// Validate the request
	violations := validateListTokensRequest(req)
	if len(violations) > 0 {
		return nil, invalidArgumentErrors(violations)
	}

	tokens, err := server.store.ListTokens(ctx, db.ListTokensParams{
		Limit:  req.GetLimit(),
		Offset: req.GetOffset(),
	})
	if err != nil {
		// Handle database errors
		return nil, handleDatabaseError(err)
	}

	rsp := &pb.ListTokensResponse{
		Tokens: convertTokenList(tokens),
	}

	return rsp, nil
}

func validateListTokensRequest(req *pb.ListTokensRequest) (violations []*CustomError) {
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
