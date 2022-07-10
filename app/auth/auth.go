package auth

import (
	"context"
	"time"
	authpb "tmpl-go-vercel/app/auth/proto/gen/go"
)

// Service implements auth service.
type Service struct {
	authpb.UnimplementedAuthServer
}

// Login logs a user in.
func (s *Service) Login(c context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	return &authpb.LoginResponse{
		AccessToken: "tkn",
		ExpiresIn:   int32(time.Second),
	}, nil
}
