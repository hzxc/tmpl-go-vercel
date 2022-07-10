package main

import (
	"log"
	"net"
	"tmpl-go-vercel/app/pingpong"
	pingpongpb "tmpl-go-vercel/gen/go/pingpong/v1"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pingpongpb.RegisterPingPongServiceServer(s, &pingpong.Service{})
	s.Serve(lis)
}
