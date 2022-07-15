package hello

import (
	"context"
	"fmt"
	"tmpl-go-vercel/app/global"
	"tmpl-go-vercel/app/grpc"

	grpc_auth "tmpl-go-vercel/app/grpc/auth"
	"tmpl-go-vercel/app/shared/auth/token"
	proto "tmpl-go-vercel/gen/go/api/hello/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func init() {
	grpc.GRPCEndpoints.Register(proto.RegisterHelloServiceServer, &Service{})
	grpc.GRPCAuthList = append(grpc.GRPCAuthList,
		fmt.Sprintf("/%s/Me", proto.HelloService_ServiceDesc.ServiceName))
}

type Service struct {
	proto.UnimplementedHelloServiceServer
}

func (s *Service) Intro(ctx context.Context, req *proto.IntroRequest) (*proto.IntroResponse, error) {

	return &proto.IntroResponse{
		Intro: fmt.Sprintf("hello, %s!", req.Name),
	}, nil
}

func (s *Service) Me(ctx context.Context, req *proto.MeRequest) (*proto.MeResponse, error) {
	aid, err := grpc_auth.AccountIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	tkn, err := grpc_auth.TokenFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	verifier := &token.JWTTokenVerifier{
		PublicKey: global.PubKey,
	}

	expiresAt, err := verifier.TokenExpiresAt(tkn)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	return &proto.MeResponse{
		Id:        1,
		Name:      aid.String(),
		Token:     tkn,
		ExpiresAt: expiresAt,
	}, nil
}
