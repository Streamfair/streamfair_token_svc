package gapi

import (
	"context"

	db "github.com/Streamfair/streamfair_token_svc/db/sqlc"
	pb "github.com/Streamfair/streamfair_token_svc/pb/refresh_token"
	"github.com/Streamfair/streamfair_token_svc/validator"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/grpc/codes"
)

func (server *Server) UpdateRefreshToken(ctx context.Context, req *pb.UpdateRefreshTokenRequest) (*pb.UpdateRefreshTokenResponse, error) {
	// Validate the request
	violations := validateUpdateRefreshTokenRequest(req)
	if len(violations) > 0 {
		return nil, invalidArgumentErrors(violations)
	}

	// Update the token in the database
	dbRefreshToken, err := server.store.UpdateRefreshToken(ctx, db.UpdateRefreshTokenParams{
		UserID: pgtype.Int8{Int64: req.GetUserId(), Valid: req.UserId != nil},
		ID:     req.GetId(),
	})
	if err != nil {
		return nil, handleDatabaseError(err)
	}

	rsp := &pb.UpdateRefreshTokenResponse{
		RefreshToken: convertRefreshToken(dbRefreshToken),
	}

	return rsp, nil
}

// validateUpdateRefreshTokenRequest validates the create token request and returns a slice of custom errors.
func validateUpdateRefreshTokenRequest(req *pb.UpdateRefreshTokenRequest) (violations []*CustomError) {
	if err := validator.ValidateId(req.GetId()); err != nil {
		violations = append(violations, (&CustomError{
			StatusCode: codes.NotFound,
		}).WithDetails("user_id", err))
	}

	if req.GetUserId() != 0 {
		if err := validator.ValidateId(req.GetUserId()); err != nil {
			violations = append(violations, (&CustomError{
				StatusCode: codes.NotFound,
			}).WithDetails("user_id", err))
		}
	}

	return violations
}
