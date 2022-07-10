package grpc

import (
	"net/http"
	"tmpl-go-vercel/app/pingpong"
	pingpongpb "tmpl-go-vercel/app/pingpong/proto/gen/go"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
)

type Handler http.Handler

func New() Handler {
	s := grpc.NewServer()
	pingpongpb.RegisterPingPongServer(s, &pingpong.Service{})
	return grpcweb.WrapServer(s, grpcweb.WithOriginFunc(func(origin string) bool {
		// Allow all origins, DO NOT do this in production
		return true
	}))

	// return s
}
