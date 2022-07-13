package hello

import (
	"context"
	"fmt"
	proto "tmpl-go-vercel/gen/go/api/hello/v1"

	"google.golang.org/grpc"
)

// func init() {
// 	grpc.GRPCEndpoints.Register(proto.RegisterHelloServiceServer, &Service{})
// }

type Service struct {
	proto.UnimplementedHelloServiceServer
}

func Register(s *grpc.Server) {
	proto.RegisterHelloServiceServer(s, &Service{})
}

func (s *Service) Intro(ctx context.Context, req *proto.IntroRequest) (*proto.IntroResponse, error) {
	return &proto.IntroResponse{
		Intro: fmt.Sprintf("hello, %s!", req.Name),
	}, nil
}
