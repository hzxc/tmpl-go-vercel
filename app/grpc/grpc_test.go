package grpc

import (
	"context"
	"log"
	"testing"
	pingpongpb "tmpl-go-vercel/app/pingpong/proto/gen/go"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Test_GrpcClient(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	conn, err := grpc.Dial("localhost:3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("can not connect server: %v", err)
	}
	cli := pingpongpb.NewPingPongClient(conn)
	resp, err := cli.PingPong(context.Background(), &pingpongpb.PingRequest{})

	if err != nil {
		log.Fatalf("can not call Ping: %v", err)
	}

	log.Println(resp)
}
