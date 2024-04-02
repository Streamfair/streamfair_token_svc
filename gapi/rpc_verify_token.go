package gapi

import (
	"context"
	"errors"

	pb "github.com/Streamfair/streamfair_token_svc/common_proto/TokenService/pb/token"
	"google.golang.org/grpc/codes"
)

func (server *Server) VerifyToken(ctx context.Context, req *pb.VerifyTokenRequest) (*pb.VerifyTokenResponse, error) {
	tokenParam := req.GetToken()

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
	token, err := server.store.GetTokenByValue(ctx, tokenParam)
	if err != nil {
		return nil, handleDatabaseError(err)
	}

	if token.Revoked {
		violation := (&CustomError{
			StatusCode: codes.InvalidArgument,
		}).WithDetails("token", errors.New("token is revoked"))
		return nil, invalidArgumentError(violation)
	}

	rsp := &pb.VerifyTokenResponse{
		Token: convertToken(token),
	}

	return rsp, nil
}
