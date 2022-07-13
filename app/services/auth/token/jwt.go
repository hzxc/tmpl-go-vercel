package token

import (
	"crypto/rsa"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// JWTTokenGen generates a JWT token.
type JWTTokenGen struct {
	privateKey *rsa.PrivateKey
	issuer     string
	nowFunc    func() time.Time
}

// NewJWTTokenGen creates a JWTTokenGen.
func NewJWTTokenGen(issuer string, privateKey *rsa.PrivateKey) *JWTTokenGen {
	return &JWTTokenGen{
		issuer:     issuer,
		nowFunc:    time.Now,
		privateKey: privateKey,
	}
}

// GenerateToken generates a token.
func (t *JWTTokenGen) GenerateToken(sub string, expire time.Duration) (string, error) {
	t.nowFunc()
	tkn := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.RegisteredClaims{
		Issuer:    t.issuer,
		IssuedAt:  jwt.NewNumericDate(t.nowFunc()),
		NotBefore: jwt.NewNumericDate(t.nowFunc()),
		ExpiresAt: jwt.NewNumericDate(t.nowFunc().Add(expire)),
		Subject:   sub,
	})

	return tkn.SignedString(t.privateKey)
}
