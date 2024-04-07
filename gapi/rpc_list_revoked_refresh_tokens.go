package gapi

import (
	"context"

	db "github.com/Streamfair/streamfair_token_svc/db/sqlc"
	pb "github.com/Streamfair/common_proto/TokenService/pb/refresh_token"
	"github.com/Streamfair/streamfair_token_svc/validator"
	"google.golang.org/grpc/codes"
)

func (server *Server) ListRevokedRefreshTokens(ctx context.Context, req *pb.ListRevokedRefreshTokensRequest) (*pb.ListRevokedRefreshTokensResponse, error) {
	// Validate the request
	violations := validateListRevokedRefreshTokensRequest(req)
	if len(violations) > 0 {
		return nil, invalidArgumentErrors(violations)
	}

	refreshTokens, err := server.store.ListRevokedRefreshTokens(ctx, db.ListRevokedRefreshTokensParams{
		Limit:  req.GetLimit(),
		Offset: req.GetOffset(),
	})
	if err != nil {
		// Handle database errors
		return nil, handleDatabaseError(err)
	}

	rsp := &pb.ListRevokedRefreshTokensResponse{
		RefreshTokens: convertRefreshTokenList(refreshTokens),
	}

	return rsp, nil
}

func validateListRevokedRefreshTokensRequest(req *pb.ListRevokedRefreshTokensRequest) (violations []*CustomError) {
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
