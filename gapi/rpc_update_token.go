package gapi

import (
	"context"

	db "github.com/Streamfair/streamfair_token_svc/db/sqlc"
	pb "github.com/Streamfair/streamfair_token_svc/pb/token"
	"github.com/Streamfair/streamfair_token_svc/validator"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/grpc/codes"
)

func (server *Server) UpdateToken(ctx context.Context, req *pb.UpdateTokenRequest) (*pb.UpdateTokenResponse, error) {
	// Validate the request
	violations := validateUpdateTokenRequest(req)
	if len(violations) > 0 {
		return nil, invalidArgumentErrors(violations)
	}

	// Update the token in the database
	dbToken, err := server.store.UpdateToken(ctx, db.UpdateTokenParams{
		UserID: pgtype.Int8{Int64: req.GetUserId(), Valid: req.UserId != nil},
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

// validateUpdateTokenRequest validates the create token request and returns a slice of custom errors.
func validateUpdateTokenRequest(req *pb.UpdateTokenRequest) (violations []*CustomError) {
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
