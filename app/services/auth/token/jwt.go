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
func (t *JWTTokenGen) GenerateToken(sub string, expire time.Duration) (string, int64, error) {
	t.nowFunc()
	expiresAt := jwt.NewNumericDate(t.nowFunc().Add(expire))
	tkn := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.RegisteredClaims{
		Issuer:    t.issuer,
		IssuedAt:  jwt.NewNumericDate(t.nowFunc()),
		NotBefore: jwt.NewNumericDate(t.nowFunc()),
		ExpiresAt: expiresAt,
		Subject:   sub,
	})

	sign, err := tkn.SignedString(t.privateKey)
	if err != nil {
		return "", 0, err
	}

	return sign, expiresAt.Unix(), nil
}
