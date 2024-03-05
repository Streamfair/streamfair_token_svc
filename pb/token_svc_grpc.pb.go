// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	refresh_token "github.com/Streamfair/streamfair_token_svc/pb/refresh_token"
	token "github.com/Streamfair/streamfair_token_svc/pb/token"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TokenServiceClient is the client API for TokenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TokenServiceClient interface {
	// Tokens
	CreateToken(ctx context.Context, in *token.CreateTokenRequest, opts ...grpc.CallOption) (*token.CreateTokenResponse, error)
	GetTokenById(ctx context.Context, in *token.GetTokenByIdRequest, opts ...grpc.CallOption) (*token.GetTokenByIdResponse, error)
	GetTokenByValue(ctx context.Context, in *token.GetTokenByValueRequest, opts ...grpc.CallOption) (*token.GetTokenByValueResponse, error)
	RevokeTokenByValue(ctx context.Context, in *token.RevokeTokenByValueRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RevokeTokenById(ctx context.Context, in *token.RevokeTokenByIdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteTokenById(ctx context.Context, in *token.DeleteTokenByIdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteTokenByValue(ctx context.Context, in *token.DeleteTokenByValueRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListRevokedTokens(ctx context.Context, in *token.ListRevokedTokensRequest, opts ...grpc.CallOption) (*token.ListRevokedTokensResponse, error)
	ListTokens(ctx context.Context, in *token.ListTokensRequest, opts ...grpc.CallOption) (*token.ListTokensResponse, error)
	UpdateToken(ctx context.Context, in *token.UpdateTokenRequest, opts ...grpc.CallOption) (*token.UpdateTokenResponse, error)
	VerifyToken(ctx context.Context, in *token.VerifyTokenRequest, opts ...grpc.CallOption) (*token.VerifyTokenResponse, error)
	// Refresh tokens
	CreateRefreshToken(ctx context.Context, in *refresh_token.CreateRefreshTokenRequest, opts ...grpc.CallOption) (*refresh_token.CreateRefreshTokenResponse, error)
	GetRefreshTokenById(ctx context.Context, in *refresh_token.GetRefreshTokenByIdRequest, opts ...grpc.CallOption) (*refresh_token.GetRefreshTokenByIdResponse, error)
	GetRefreshTokenByValue(ctx context.Context, in *refresh_token.GetRefreshTokenByValueRequest, opts ...grpc.CallOption) (*refresh_token.GetRefreshTokenByValueResponse, error)
	RevokeRefreshTokenByValue(ctx context.Context, in *refresh_token.RevokeRefreshTokenByValueRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RevokeRefreshTokenById(ctx context.Context, in *refresh_token.RevokeRefreshTokenByIdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteRefreshTokenById(ctx context.Context, in *refresh_token.DeleteRefreshTokenByIdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteRefreshTokenByValue(ctx context.Context, in *refresh_token.DeleteRefreshTokenByValueRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListRevokedRefreshTokens(ctx context.Context, in *refresh_token.ListRevokedRefreshTokensRequest, opts ...grpc.CallOption) (*refresh_token.ListRevokedRefreshTokensResponse, error)
	ListRefreshTokens(ctx context.Context, in *refresh_token.ListRefreshTokensRequest, opts ...grpc.CallOption) (*refresh_token.ListRefreshTokensResponse, error)
	UpdateRefreshToken(ctx context.Context, in *refresh_token.UpdateRefreshTokenRequest, opts ...grpc.CallOption) (*refresh_token.UpdateRefreshTokenResponse, error)
	VerifyRefreshToken(ctx context.Context, in *refresh_token.VerifyRefreshTokenRequest, opts ...grpc.CallOption) (*refresh_token.VerifyRefreshTokenResponse, error)
}

type tokenServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTokenServiceClient(cc grpc.ClientConnInterface) TokenServiceClient {
	return &tokenServiceClient{cc}
}

func (c *tokenServiceClient) CreateToken(ctx context.Context, in *token.CreateTokenRequest, opts ...grpc.CallOption) (*token.CreateTokenResponse, error) {
	out := new(token.CreateTokenResponse)
	err := c.cc.Invoke(ctx, "/pb.TokenService/CreateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) GetTokenById(ctx context.Context, in *token.GetTokenByIdRequest, opts ...grpc.CallOption) (*token.GetTokenByIdResponse, error) {
	out := new(token.GetTokenByIdResponse)
	err := c.cc.Invoke(ctx, "/pb.TokenService/GetTokenById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) GetTokenByValue(ctx context.Context, in *token.GetTokenByValueRequest, opts ...grpc.CallOption) (*token.GetTokenByValueResponse, error) {
	out := new(token.GetTokenByValueResponse)
	err := c.cc.Invoke(ctx, "/pb.TokenService/GetTokenByValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) RevokeTokenByValue(ctx context.Context, in *token.RevokeTokenByValueRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.TokenService/RevokeTokenByValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) RevokeTokenById(ctx context.Context, in *token.RevokeTokenByIdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.TokenService/RevokeTokenById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) DeleteTokenById(ctx context.Context, in *token.DeleteTokenByIdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.TokenService/DeleteTokenById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) DeleteTokenByValue(ctx context.Context, in *token.DeleteTokenByValueRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.TokenService/DeleteTokenByValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) ListRevokedTokens(ctx context.Context, in *token.ListRevokedTokensRequest, opts ...grpc.CallOption) (*token.ListRevokedTokensResponse, error) {
	out := new(token.ListRevokedTokensResponse)
	err := c.cc.Invoke(ctx, "/pb.TokenService/ListRevokedTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) ListTokens(ctx context.Context, in *token.ListTokensRequest, opts ...grpc.CallOption) (*token.ListTokensResponse, error) {
	out := new(token.ListTokensResponse)
	err := c.cc.Invoke(ctx, "/pb.TokenService/ListTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) UpdateToken(ctx context.Context, in *token.UpdateTokenRequest, opts ...grpc.CallOption) (*token.UpdateTokenResponse, error) {
	out := new(token.UpdateTokenResponse)
	err := c.cc.Invoke(ctx, "/pb.TokenService/UpdateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) VerifyToken(ctx context.Context, in *token.VerifyTokenRequest, opts ...grpc.CallOption) (*token.VerifyTokenResponse, error) {
	out := new(token.VerifyTokenResponse)
	err := c.cc.Invoke(ctx, "/pb.TokenService/VerifyToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) CreateRefreshToken(ctx context.Context, in *refresh_token.CreateRefreshTokenRequest, opts ...grpc.CallOption) (*refresh_token.CreateRefreshTokenResponse, error) {
	out := new(refresh_token.CreateRefreshTokenResponse)
	err := c.cc.Invoke(ctx, "/pb.TokenService/CreateRefreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) GetRefreshTokenById(ctx context.Context, in *refresh_token.GetRefreshTokenByIdRequest, opts ...grpc.CallOption) (*refresh_token.GetRefreshTokenByIdResponse, error) {
	out := new(refresh_token.GetRefreshTokenByIdResponse)
	err := c.cc.Invoke(ctx, "/pb.TokenService/GetRefreshTokenById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) GetRefreshTokenByValue(ctx context.Context, in *refresh_token.GetRefreshTokenByValueRequest, opts ...grpc.CallOption) (*refresh_token.GetRefreshTokenByValueResponse, error) {
	out := new(refresh_token.GetRefreshTokenByValueResponse)
	err := c.cc.Invoke(ctx, "/pb.TokenService/GetRefreshTokenByValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) RevokeRefreshTokenByValue(ctx context.Context, in *refresh_token.RevokeRefreshTokenByValueRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.TokenService/RevokeRefreshTokenByValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) RevokeRefreshTokenById(ctx context.Context, in *refresh_token.RevokeRefreshTokenByIdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.TokenService/RevokeRefreshTokenById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) DeleteRefreshTokenById(ctx context.Context, in *refresh_token.DeleteRefreshTokenByIdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.TokenService/DeleteRefreshTokenById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) DeleteRefreshTokenByValue(ctx context.Context, in *refresh_token.DeleteRefreshTokenByValueRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.TokenService/DeleteRefreshTokenByValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) ListRevokedRefreshTokens(ctx context.Context, in *refresh_token.ListRevokedRefreshTokensRequest, opts ...grpc.CallOption) (*refresh_token.ListRevokedRefreshTokensResponse, error) {
	out := new(refresh_token.ListRevokedRefreshTokensResponse)
	err := c.cc.Invoke(ctx, "/pb.TokenService/ListRevokedRefreshTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) ListRefreshTokens(ctx context.Context, in *refresh_token.ListRefreshTokensRequest, opts ...grpc.CallOption) (*refresh_token.ListRefreshTokensResponse, error) {
	out := new(refresh_token.ListRefreshTokensResponse)
	err := c.cc.Invoke(ctx, "/pb.TokenService/ListRefreshTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) UpdateRefreshToken(ctx context.Context, in *refresh_token.UpdateRefreshTokenRequest, opts ...grpc.CallOption) (*refresh_token.UpdateRefreshTokenResponse, error) {
	out := new(refresh_token.UpdateRefreshTokenResponse)
	err := c.cc.Invoke(ctx, "/pb.TokenService/UpdateRefreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) VerifyRefreshToken(ctx context.Context, in *refresh_token.VerifyRefreshTokenRequest, opts ...grpc.CallOption) (*refresh_token.VerifyRefreshTokenResponse, error) {
	out := new(refresh_token.VerifyRefreshTokenResponse)
	err := c.cc.Invoke(ctx, "/pb.TokenService/VerifyRefreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TokenServiceServer is the server API for TokenService service.
// All implementations must embed UnimplementedTokenServiceServer
// for forward compatibility
type TokenServiceServer interface {
	// Tokens
	CreateToken(context.Context, *token.CreateTokenRequest) (*token.CreateTokenResponse, error)
	GetTokenById(context.Context, *token.GetTokenByIdRequest) (*token.GetTokenByIdResponse, error)
	GetTokenByValue(context.Context, *token.GetTokenByValueRequest) (*token.GetTokenByValueResponse, error)
	RevokeTokenByValue(context.Context, *token.RevokeTokenByValueRequest) (*emptypb.Empty, error)
	RevokeTokenById(context.Context, *token.RevokeTokenByIdRequest) (*emptypb.Empty, error)
	DeleteTokenById(context.Context, *token.DeleteTokenByIdRequest) (*emptypb.Empty, error)
	DeleteTokenByValue(context.Context, *token.DeleteTokenByValueRequest) (*emptypb.Empty, error)
	ListRevokedTokens(context.Context, *token.ListRevokedTokensRequest) (*token.ListRevokedTokensResponse, error)
	ListTokens(context.Context, *token.ListTokensRequest) (*token.ListTokensResponse, error)
	UpdateToken(context.Context, *token.UpdateTokenRequest) (*token.UpdateTokenResponse, error)
	VerifyToken(context.Context, *token.VerifyTokenRequest) (*token.VerifyTokenResponse, error)
	// Refresh tokens
	CreateRefreshToken(context.Context, *refresh_token.CreateRefreshTokenRequest) (*refresh_token.CreateRefreshTokenResponse, error)
	GetRefreshTokenById(context.Context, *refresh_token.GetRefreshTokenByIdRequest) (*refresh_token.GetRefreshTokenByIdResponse, error)
	GetRefreshTokenByValue(context.Context, *refresh_token.GetRefreshTokenByValueRequest) (*refresh_token.GetRefreshTokenByValueResponse, error)
	RevokeRefreshTokenByValue(context.Context, *refresh_token.RevokeRefreshTokenByValueRequest) (*emptypb.Empty, error)
	RevokeRefreshTokenById(context.Context, *refresh_token.RevokeRefreshTokenByIdRequest) (*emptypb.Empty, error)
	DeleteRefreshTokenById(context.Context, *refresh_token.DeleteRefreshTokenByIdRequest) (*emptypb.Empty, error)
	DeleteRefreshTokenByValue(context.Context, *refresh_token.DeleteRefreshTokenByValueRequest) (*emptypb.Empty, error)
	ListRevokedRefreshTokens(context.Context, *refresh_token.ListRevokedRefreshTokensRequest) (*refresh_token.ListRevokedRefreshTokensResponse, error)
	ListRefreshTokens(context.Context, *refresh_token.ListRefreshTokensRequest) (*refresh_token.ListRefreshTokensResponse, error)
	UpdateRefreshToken(context.Context, *refresh_token.UpdateRefreshTokenRequest) (*refresh_token.UpdateRefreshTokenResponse, error)
	VerifyRefreshToken(context.Context, *refresh_token.VerifyRefreshTokenRequest) (*refresh_token.VerifyRefreshTokenResponse, error)
	mustEmbedUnimplementedTokenServiceServer()
}

// UnimplementedTokenServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTokenServiceServer struct {
}

func (UnimplementedTokenServiceServer) CreateToken(context.Context, *token.CreateTokenRequest) (*token.CreateTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateToken not implemented")
}
func (UnimplementedTokenServiceServer) GetTokenById(context.Context, *token.GetTokenByIdRequest) (*token.GetTokenByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTokenById not implemented")
}
func (UnimplementedTokenServiceServer) GetTokenByValue(context.Context, *token.GetTokenByValueRequest) (*token.GetTokenByValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTokenByValue not implemented")
}
func (UnimplementedTokenServiceServer) RevokeTokenByValue(context.Context, *token.RevokeTokenByValueRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeTokenByValue not implemented")
}
func (UnimplementedTokenServiceServer) RevokeTokenById(context.Context, *token.RevokeTokenByIdRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeTokenById not implemented")
}
func (UnimplementedTokenServiceServer) DeleteTokenById(context.Context, *token.DeleteTokenByIdRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTokenById not implemented")
}
func (UnimplementedTokenServiceServer) DeleteTokenByValue(context.Context, *token.DeleteTokenByValueRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTokenByValue not implemented")
}
func (UnimplementedTokenServiceServer) ListRevokedTokens(context.Context, *token.ListRevokedTokensRequest) (*token.ListRevokedTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRevokedTokens not implemented")
}
func (UnimplementedTokenServiceServer) ListTokens(context.Context, *token.ListTokensRequest) (*token.ListTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTokens not implemented")
}
func (UnimplementedTokenServiceServer) UpdateToken(context.Context, *token.UpdateTokenRequest) (*token.UpdateTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateToken not implemented")
}
func (UnimplementedTokenServiceServer) VerifyToken(context.Context, *token.VerifyTokenRequest) (*token.VerifyTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyToken not implemented")
}
func (UnimplementedTokenServiceServer) CreateRefreshToken(context.Context, *refresh_token.CreateRefreshTokenRequest) (*refresh_token.CreateRefreshTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRefreshToken not implemented")
}
func (UnimplementedTokenServiceServer) GetRefreshTokenById(context.Context, *refresh_token.GetRefreshTokenByIdRequest) (*refresh_token.GetRefreshTokenByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRefreshTokenById not implemented")
}
func (UnimplementedTokenServiceServer) GetRefreshTokenByValue(context.Context, *refresh_token.GetRefreshTokenByValueRequest) (*refresh_token.GetRefreshTokenByValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRefreshTokenByValue not implemented")
}
func (UnimplementedTokenServiceServer) RevokeRefreshTokenByValue(context.Context, *refresh_token.RevokeRefreshTokenByValueRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeRefreshTokenByValue not implemented")
}
func (UnimplementedTokenServiceServer) RevokeRefreshTokenById(context.Context, *refresh_token.RevokeRefreshTokenByIdRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeRefreshTokenById not implemented")
}
func (UnimplementedTokenServiceServer) DeleteRefreshTokenById(context.Context, *refresh_token.DeleteRefreshTokenByIdRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRefreshTokenById not implemented")
}
func (UnimplementedTokenServiceServer) DeleteRefreshTokenByValue(context.Context, *refresh_token.DeleteRefreshTokenByValueRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRefreshTokenByValue not implemented")
}
func (UnimplementedTokenServiceServer) ListRevokedRefreshTokens(context.Context, *refresh_token.ListRevokedRefreshTokensRequest) (*refresh_token.ListRevokedRefreshTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRevokedRefreshTokens not implemented")
}
func (UnimplementedTokenServiceServer) ListRefreshTokens(context.Context, *refresh_token.ListRefreshTokensRequest) (*refresh_token.ListRefreshTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRefreshTokens not implemented")
}
func (UnimplementedTokenServiceServer) UpdateRefreshToken(context.Context, *refresh_token.UpdateRefreshTokenRequest) (*refresh_token.UpdateRefreshTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRefreshToken not implemented")
}
func (UnimplementedTokenServiceServer) VerifyRefreshToken(context.Context, *refresh_token.VerifyRefreshTokenRequest) (*refresh_token.VerifyRefreshTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyRefreshToken not implemented")
}
func (UnimplementedTokenServiceServer) mustEmbedUnimplementedTokenServiceServer() {}

// UnsafeTokenServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TokenServiceServer will
// result in compilation errors.
type UnsafeTokenServiceServer interface {
	mustEmbedUnimplementedTokenServiceServer()
}

func RegisterTokenServiceServer(s grpc.ServiceRegistrar, srv TokenServiceServer) {
	s.RegisterService(&TokenService_ServiceDesc, srv)
}

func _TokenService_CreateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(token.CreateTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).CreateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TokenService/CreateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).CreateToken(ctx, req.(*token.CreateTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_GetTokenById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(token.GetTokenByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).GetTokenById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TokenService/GetTokenById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).GetTokenById(ctx, req.(*token.GetTokenByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_GetTokenByValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(token.GetTokenByValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).GetTokenByValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TokenService/GetTokenByValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).GetTokenByValue(ctx, req.(*token.GetTokenByValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_RevokeTokenByValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(token.RevokeTokenByValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).RevokeTokenByValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TokenService/RevokeTokenByValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).RevokeTokenByValue(ctx, req.(*token.RevokeTokenByValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_RevokeTokenById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(token.RevokeTokenByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).RevokeTokenById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TokenService/RevokeTokenById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).RevokeTokenById(ctx, req.(*token.RevokeTokenByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_DeleteTokenById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(token.DeleteTokenByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).DeleteTokenById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TokenService/DeleteTokenById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).DeleteTokenById(ctx, req.(*token.DeleteTokenByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_DeleteTokenByValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(token.DeleteTokenByValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).DeleteTokenByValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TokenService/DeleteTokenByValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).DeleteTokenByValue(ctx, req.(*token.DeleteTokenByValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_ListRevokedTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(token.ListRevokedTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).ListRevokedTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TokenService/ListRevokedTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).ListRevokedTokens(ctx, req.(*token.ListRevokedTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_ListTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(token.ListTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).ListTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TokenService/ListTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).ListTokens(ctx, req.(*token.ListTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_UpdateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(token.UpdateTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).UpdateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TokenService/UpdateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).UpdateToken(ctx, req.(*token.UpdateTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_VerifyToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(token.VerifyTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).VerifyToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TokenService/VerifyToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).VerifyToken(ctx, req.(*token.VerifyTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_CreateRefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(refresh_token.CreateRefreshTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).CreateRefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TokenService/CreateRefreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).CreateRefreshToken(ctx, req.(*refresh_token.CreateRefreshTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_GetRefreshTokenById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(refresh_token.GetRefreshTokenByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).GetRefreshTokenById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TokenService/GetRefreshTokenById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).GetRefreshTokenById(ctx, req.(*refresh_token.GetRefreshTokenByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_GetRefreshTokenByValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(refresh_token.GetRefreshTokenByValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).GetRefreshTokenByValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TokenService/GetRefreshTokenByValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).GetRefreshTokenByValue(ctx, req.(*refresh_token.GetRefreshTokenByValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_RevokeRefreshTokenByValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(refresh_token.RevokeRefreshTokenByValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).RevokeRefreshTokenByValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TokenService/RevokeRefreshTokenByValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).RevokeRefreshTokenByValue(ctx, req.(*refresh_token.RevokeRefreshTokenByValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_RevokeRefreshTokenById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(refresh_token.RevokeRefreshTokenByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).RevokeRefreshTokenById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TokenService/RevokeRefreshTokenById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).RevokeRefreshTokenById(ctx, req.(*refresh_token.RevokeRefreshTokenByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_DeleteRefreshTokenById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(refresh_token.DeleteRefreshTokenByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).DeleteRefreshTokenById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TokenService/DeleteRefreshTokenById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).DeleteRefreshTokenById(ctx, req.(*refresh_token.DeleteRefreshTokenByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_DeleteRefreshTokenByValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(refresh_token.DeleteRefreshTokenByValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).DeleteRefreshTokenByValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TokenService/DeleteRefreshTokenByValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).DeleteRefreshTokenByValue(ctx, req.(*refresh_token.DeleteRefreshTokenByValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_ListRevokedRefreshTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(refresh_token.ListRevokedRefreshTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).ListRevokedRefreshTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TokenService/ListRevokedRefreshTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).ListRevokedRefreshTokens(ctx, req.(*refresh_token.ListRevokedRefreshTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_ListRefreshTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(refresh_token.ListRefreshTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).ListRefreshTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TokenService/ListRefreshTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).ListRefreshTokens(ctx, req.(*refresh_token.ListRefreshTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_UpdateRefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(refresh_token.UpdateRefreshTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).UpdateRefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TokenService/UpdateRefreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).UpdateRefreshToken(ctx, req.(*refresh_token.UpdateRefreshTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_VerifyRefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(refresh_token.VerifyRefreshTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).VerifyRefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TokenService/VerifyRefreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).VerifyRefreshToken(ctx, req.(*refresh_token.VerifyRefreshTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TokenService_ServiceDesc is the grpc.ServiceDesc for TokenService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TokenService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.TokenService",
	HandlerType: (*TokenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateToken",
			Handler:    _TokenService_CreateToken_Handler,
		},
		{
			MethodName: "GetTokenById",
			Handler:    _TokenService_GetTokenById_Handler,
		},
		{
			MethodName: "GetTokenByValue",
			Handler:    _TokenService_GetTokenByValue_Handler,
		},
		{
			MethodName: "RevokeTokenByValue",
			Handler:    _TokenService_RevokeTokenByValue_Handler,
		},
		{
			MethodName: "RevokeTokenById",
			Handler:    _TokenService_RevokeTokenById_Handler,
		},
		{
			MethodName: "DeleteTokenById",
			Handler:    _TokenService_DeleteTokenById_Handler,
		},
		{
			MethodName: "DeleteTokenByValue",
			Handler:    _TokenService_DeleteTokenByValue_Handler,
		},
		{
			MethodName: "ListRevokedTokens",
			Handler:    _TokenService_ListRevokedTokens_Handler,
		},
		{
			MethodName: "ListTokens",
			Handler:    _TokenService_ListTokens_Handler,
		},
		{
			MethodName: "UpdateToken",
			Handler:    _TokenService_UpdateToken_Handler,
		},
		{
			MethodName: "VerifyToken",
			Handler:    _TokenService_VerifyToken_Handler,
		},
		{
			MethodName: "CreateRefreshToken",
			Handler:    _TokenService_CreateRefreshToken_Handler,
		},
		{
			MethodName: "GetRefreshTokenById",
			Handler:    _TokenService_GetRefreshTokenById_Handler,
		},
		{
			MethodName: "GetRefreshTokenByValue",
			Handler:    _TokenService_GetRefreshTokenByValue_Handler,
		},
		{
			MethodName: "RevokeRefreshTokenByValue",
			Handler:    _TokenService_RevokeRefreshTokenByValue_Handler,
		},
		{
			MethodName: "RevokeRefreshTokenById",
			Handler:    _TokenService_RevokeRefreshTokenById_Handler,
		},
		{
			MethodName: "DeleteRefreshTokenById",
			Handler:    _TokenService_DeleteRefreshTokenById_Handler,
		},
		{
			MethodName: "DeleteRefreshTokenByValue",
			Handler:    _TokenService_DeleteRefreshTokenByValue_Handler,
		},
		{
			MethodName: "ListRevokedRefreshTokens",
			Handler:    _TokenService_ListRevokedRefreshTokens_Handler,
		},
		{
			MethodName: "ListRefreshTokens",
			Handler:    _TokenService_ListRefreshTokens_Handler,
		},
		{
			MethodName: "UpdateRefreshToken",
			Handler:    _TokenService_UpdateRefreshToken_Handler,
		},
		{
			MethodName: "VerifyRefreshToken",
			Handler:    _TokenService_VerifyRefreshToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "token_svc.proto",
}
