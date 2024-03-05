package gapi

import (
	"context"

	db "github.com/Streamfair/streamfair_token_svc/db/sqlc"
	pb "github.com/Streamfair/streamfair_token_svc/pb/refresh_token"
	"github.com/Streamfair/streamfair_token_svc/validator"
	"google.golang.org/grpc/codes"
)

func (server *Server) ListRefreshTokens(ctx context.Context, req *pb.ListRefreshTokensRequest) (*pb.ListRefreshTokensResponse, error) {
	// Validate the request
	violations := validateListRefreshTokensRequest(req)
	if len(violations) > 0 {
		return nil, invalidArgumentErrors(violations)
	}

	refreshTokens, err := server.store.ListRefreshTokens(ctx, db.ListRefreshTokensParams{
		Limit:  req.GetLimit(),
		Offset: req.GetOffset(),
	})
	if err != nil {
		// Handle database errors
		return nil, handleDatabaseError(err)
	}

	rsp := &pb.ListRefreshTokensResponse{
		RefreshTokens: convertRefreshTokenList(refreshTokens),
	}

	return rsp, nil
}

func validateListRefreshTokensRequest(req *pb.ListRefreshTokensRequest) (violations []*CustomError) {
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
