package gapi

import (
	"context"
	"errors"

	pb "github.com/Streamfair/streamfair_token_svc/pb/refresh_token"
	"google.golang.org/grpc/codes"
)

func (server *Server) VerifyRefreshToken(ctx context.Context, req *pb.VerifyRefreshTokenRequest) (*pb.VerifyRefreshTokenResponse, error) {
	tokenParam := req.GetRefreshToken()

	// Perform field validation
	if tokenParam == "" {
		violation := (&CustomError{
			StatusCode: codes.InvalidArgument,
		}).WithDetails("token", errors.New("token is required"))
		return nil, invalidArgumentError(violation)
	}

	_, err := server.localTokenMaker.VerifyLocalToken(tokenParam)
	if err != nil {
		violation := (&CustomError{
			StatusCode: codes.InvalidArgument,
		}).WithDetails("token", errors.New("token is invalid"))
		return nil, invalidArgumentError(violation)
	}

	// Check if the token exists
	refreshToken, err := server.store.GetRefreshTokenByValue(ctx, tokenParam)
	if err != nil {
		return nil, handleDatabaseError(err)
	}

	if refreshToken.Revoked {
		violation := (&CustomError{
			StatusCode: codes.InvalidArgument,
		}).WithDetails("token", errors.New("token is revoked"))
		return nil, invalidArgumentError(violation)
	}

	rsp := &pb.VerifyRefreshTokenResponse{
		RefreshToken: convertRefreshToken(refreshToken),
	}

	return rsp, nil
}
