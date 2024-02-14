package gapi

import (
	"context"
	"time"

	db "github.com/Streamfair/streamfair_token_svc/db/sqlc"
	pb "github.com/Streamfair/streamfair_token_svc/pb/token"
	"github.com/Streamfair/streamfair_token_svc/validator"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateToken(ctx context.Context, req *pb.CreateTokenRequest) (*pb.CreateTokenResponse, error) {
	// Validate the request
	violations := validateCreateTokenRequest(req)
	if len(violations) > 0 {
		return nil, invalidArgumentError(violations)
	}
	
	// Create a new token
	duration, _ := time.ParseDuration(req.GetExpiresAt())
	accessToken, accessPayload, err := server.localTokenMaker.CreateLocalToken(duration)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create access token.")
	}

	// Save the token to the database
	dbToken, err := server.store.CreateToken(ctx, db.CreateTokenParams{
		UserID:    req.GetUserId(),
		TokenType: pgtype.Text{String: "access", Valid: true},
		Token:     accessToken,
		ExpiresAt: accessPayload.ExpiredAt,
	})
	if err != nil {
		return nil, handleDatabaseError(err)
	}

	rsp := &pb.CreateTokenResponse{
		Token: convertToken(dbToken),
	}

	return rsp, nil
}

func validateCreateTokenRequest(req *pb.CreateTokenRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validator.ValidateUserId(req.GetUserId()); err != nil {
		violations = append(violations, fieldViolation("user_id", err))
	}

	if err := validator.ValidateDuration(req.GetExpiresAt()); err != nil {
		violations = append(violations, fieldViolation("expires_at", err))
	}

	return violations
}
