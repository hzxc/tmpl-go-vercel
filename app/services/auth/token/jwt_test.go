package token

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func Test_TimeUnix(t *testing.T) {
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().UnixMilli())
	fmt.Println(time.Now().UnixNano())
	fmt.Println(time.Now().UnixMicro())
}

const privteKey = `-----BEGIN RSA PRIVATE KEY-----
PRIVATE KEY PRIVATE KEY PRIVATE KEY
-----END RSA PRIVATE KEY-----`

func TestGenerateToken(t *testing.T) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privteKey))
	if err != nil {
		t.Fatalf("cannot parse private key: %v", err)
	}
	g := NewJWTTokenGen("hello/auth", key)
	g.nowFunc = func() time.Time {
		return time.Unix(1657714993, 0)
	}
	tkn, err := g.GenerateToken("username", time.Hour)
	if err != nil {
		t.Errorf("cannot generate token: %v", err)
	}

	want := "HEADER.PAYLOAD:DATA.VERIFY SIGNATURE"
	if tkn != want {
		t.Errorf("wrong token generated. want: %q; got: %q", want, tkn)
	}
}
