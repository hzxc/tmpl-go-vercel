package proto

import (
	"testing"
	authpb "tmpl-go-vercel/app/auth/proto/gen/go"
)

func Test_Sample(t *testing.T) {
	req := authpb.LoginRequest{
		Code: "Success",
	}
	t.Log(&req)
}
