package hello

import (
	"context"
	"fmt"
	"tmpl-go-vercel/app/grpc"
	proto "tmpl-go-vercel/gen/go/api/hello/v1"
)

func init() {
	grpc.GRPCEndpoints.Register(proto.RegisterHelloServiceServer, &Service{})
	grpc.GRPCAuthList = append(grpc.GRPCAuthList,
		fmt.Sprintf("/%s/Intro", proto.HelloService_ServiceDesc.ServiceName))
}

type Service struct {
	proto.UnimplementedHelloServiceServer
}

func (s *Service) Intro(ctx context.Context, req *proto.IntroRequest) (*proto.IntroResponse, error) {

	return &proto.IntroResponse{
		Intro: fmt.Sprintf("hello, %s!", req.Name),
	}, nil
}
