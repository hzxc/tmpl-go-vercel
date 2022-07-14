package grpc

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"

	proto "tmpl-go-vercel/gen/go/api/hello/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Test_GrpcClient(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	conn, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("can not connect server: %v", err)
	}
	cli := proto.NewHelloServiceClient(conn)
	resp, err := cli.Intro(context.Background(), &proto.IntroRequest{Name: "foo"})

	if err != nil {
		log.Fatalf("can not call Ping: %v", err)
	}

	log.Println(resp)
}

func Test_Str(t *testing.T) {
	str := "not found: ☀️   🌡️+47°F 🌬️↖2mph\n"
	fmt.Println(str)
	fmt.Println(strings.Trim(str, "\n"))
}
