package main

import (
	"log"
	"net"
	"tmpl-go-vercel/app/services/hello"
	proto "tmpl-go-vercel/gen/go/api/hello/v1"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterHelloServiceServer(s, &hello.Service{})
	s.Serve(lis)
}
