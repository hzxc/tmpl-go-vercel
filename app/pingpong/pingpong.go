package pingpong

import (
	"context"
	pingpongpb "tmpl-go-vercel/app/pingpong/proto/gen/go"
)

type Service struct {
	pingpongpb.UnimplementedPingPongServer
}

func (s *Service) PingPong(ctx context.Context, request *pingpongpb.PingRequest) (*pingpongpb.PongResponse, error) {
	return &pingpongpb.PongResponse{
		Pong: "pong",
	}, nil
}
