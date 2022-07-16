package project

import (
	"context"
	"fmt"
	"tmpl-go-vercel/app/grpc"

	proto "tmpl-go-vercel/gen/go/api/project/v1"
)

func init() {
	grpc.GRPCEndpoints.Register(proto.RegisterProjectServiceServer, &Service{})
	grpc.GRPCAuthList = append(grpc.GRPCAuthList,
		fmt.Sprintf("/%s/List", proto.ProjectService_ServiceDesc.ServiceName),
		fmt.Sprintf("/%s/Create", proto.ProjectService_ServiceDesc.ServiceName),
		fmt.Sprintf("/%s/Edit", proto.ProjectService_ServiceDesc.ServiceName),
		fmt.Sprintf("/%s/Delete", proto.ProjectService_ServiceDesc.ServiceName),
	)
}

type Service struct {
	proto.UnimplementedProjectServiceServer
}

func (s *Service) Create(ctx context.Context, req *proto.CreateRequest) (*proto.CreateResponse, error) {

	return &proto.CreateResponse{
		// Intro: fmt.Sprintf("hello, %s!", req.Name),
	}, nil
}
