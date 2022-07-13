package healthcheck

import (
	"context"
	proto "tmpl-go-vercel/gen/go/api/healthcheck/v1"

	v "gomodules.xyz/x/version"
	"google.golang.org/grpc"
)

type Service struct {
	proto.UnimplementedStatusServiceServer
}

func Register(s *grpc.Server) {
	proto.RegisterStatusServiceServer(s, &Service{})
}

func (s *Service) Status(ctx context.Context, req *proto.StatusRequest) (*proto.StatusResponse, error) {
	return &proto.StatusResponse{
		Version: &proto.Version{
			Version:         v.Version.Version,
			VersionStrategy: v.Version.VersionStrategy,
			CommitHash:      v.Version.CommitHash,
			GitBranch:       v.Version.GitBranch,
			GitTag:          v.Version.GitTag,
			CommitTimestamp: v.Version.CommitTimestamp,
		},
	}, nil
}
