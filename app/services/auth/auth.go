package auth

import (
	"context"
	"time"
	"tmpl-go-vercel/app/global"
	"tmpl-go-vercel/app/grpc"
	"tmpl-go-vercel/app/services/auth/token"

	proto "tmpl-go-vercel/gen/go/api/auth/v1"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func init() {

	grpc.GRPCEndpoints.Register(proto.RegisterAuthServiceServer, &Service{
		TokenExpire:    time.Hour,
		TokenGenerator: token.NewJWTTokenGen("hello/auth", global.PrivKey),
	})
}

// Service implements auth service.
type Service struct {
	TokenGenerator TokenGenerator
	TokenExpire    time.Duration
	proto.UnsafeAuthServiceServer
}

type TokenGenerator interface {
	GenerateToken(accountID string, expire time.Duration) (string, int64, error)
}

func (s *Service) Login(c context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	var (
		tkn       string
		err       error
		expiresAt int64
	)

	if global.Config.Username == req.Username && global.Config.Password == req.Password {
		tkn, expiresAt, err = s.TokenGenerator.GenerateToken(req.Username, s.TokenExpire)
		if err != nil {
			zap.L().Error("cannot generate token", zap.Error(err))
			return nil, status.Error(codes.Internal, "")
		}
	} else {
		zap.L().Warn("username or password invalid", zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, "")
	}
	return &proto.LoginResponse{
		Id:        1,
		Name:      req.Username,
		Token:     tkn,
		ExpiresAt: expiresAt,
	}, nil
}
