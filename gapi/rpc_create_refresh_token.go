package gapi

import (
	"context"
	"time"

	db "github.com/Streamfair/streamfair_token_svc/db/sqlc"
	pb "github.com/Streamfair/streamfair_token_svc/pb/refresh_token"
	"github.com/Streamfair/streamfair_token_svc/validator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateRefreshToken(ctx context.Context, req *pb.CreateRefreshTokenRequest) (*pb.CreateRefreshTokenResponse, error) {
	// Validate the request
	violations := validateCreateRefreshTokenRequest(req)
	if len(violations) > 0 {
		return nil, invalidArgumentErrors(violations)
	}

	// Create a new token
	duration, _ := time.ParseDuration(req.GetExpiresAt())
	refreshToken, accessPayload, err := server.localTokenMaker.CreateLocalToken(duration)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create access token.")
	}

	// Save the token to the database
	dbRefreshToken, err := server.store.CreateRefreshToken(ctx, db.CreateRefreshTokenParams{
		UserID:    req.GetUserId(),
		Token:     refreshToken,
		ExpiresAt: accessPayload.ExpiredAt,
	})
	if err != nil {
		return nil, handleDatabaseError(err)
	}

	rsp := &pb.CreateRefreshTokenResponse{
		RefreshToken: convertRefreshToken(dbRefreshToken),
	}

	return rsp, nil
}

// validateCreateRefreshTokenRequest validates the create token request and returns a slice of custom errors.
func validateCreateRefreshTokenRequest(req *pb.CreateRefreshTokenRequest) (violations []*CustomError) {
	if err := validator.ValidateId(req.GetUserId()); err != nil {
		violations = append(violations, (&CustomError{
			StatusCode: codes.NotFound,
		}).WithDetails("user_id", err))
	}

	if err := validator.ValidateDuration(req.GetExpiresAt()); err != nil {
		violations = append(violations, (&CustomError{
			StatusCode: codes.OutOfRange,
		}).WithDetails("expires_at", err))
	}

	return violations
}
