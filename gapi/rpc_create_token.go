package gapi

import (
	"context"
	"time"

	pb "github.com/Streamfair/common_proto/TokenService/pb/token"
	db "github.com/Streamfair/streamfair_token_svc/db/sqlc"
	"github.com/Streamfair/streamfair_token_svc/validator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) CreateToken(ctx context.Context, req *pb.CreateTokenRequest) (*pb.CreateTokenResponse, error) {
	// Validate the request
	violations := validateCreateTokenRequest(req)
	if len(violations) > 0 {
		return nil, invalidArgumentErrors(violations)
	}

	// Create a new token
	duration, _ := time.ParseDuration(req.GetExpiresAt())
	accessToken, accessPayload, err := server.localTokenMaker.CreateLocalToken(duration)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create access token")
	}

	// Save the token to the database
	dbToken, err := server.store.CreateToken(ctx, db.CreateTokenParams{
		UserID:    req.GetUserId(),
		Token:     accessToken,
		ExpiresAt: accessPayload.ExpiredAt,
	})
	if err != nil {
		return nil, handleDatabaseError(err)
	}

	rsp := &pb.CreateTokenResponse{
		Token: convertToken(dbToken),
		Payload: &pb.TokenPayload{
			Uuid:     accessPayload.ID.String(),
			IssuedAt: timestamppb.New(accessPayload.IssuedAt),
			ExpiredAt: timestamppb.New(accessPayload.ExpiredAt),
		},
	}

	return rsp, nil
}

// validateCreateTokenRequest validates the create token request and returns a slice of custom errors.
func validateCreateTokenRequest(req *pb.CreateTokenRequest) (violations []*CustomError) {
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
