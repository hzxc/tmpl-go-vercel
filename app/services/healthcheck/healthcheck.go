package healthcheck

import (
	"context"
	"tmpl-go-vercel/app/grpc"
	proto "tmpl-go-vercel/gen/go/api/healthcheck/v1"
)

func init() {
	grpc.GRPCEndpoints.Register(proto.RegisterStatusServiceServer, &Service{})
}

type Service struct {
	proto.UnimplementedStatusServiceServer
}

func (s *Service) Status(ctx context.Context, req *proto.StatusRequest) (*proto.StatusResponse, error) {
	return &proto.StatusResponse{
		Version: &proto.Version{
			Name:            "VercelGo",
			Version:         "Version",
			VersionStrategy: "VersionStrategy",
			CommitHash:      "CommitHash",
			GitBranch:       "GitBranch",
			GitTag:          "GitTag",
			CommitTimestamp: "CommitTimestamp",
		},
	}, nil
}
