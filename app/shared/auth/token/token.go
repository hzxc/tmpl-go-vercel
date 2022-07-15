package token

import (
	"crypto/rsa"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

// JWTTokenVerifier verifies jwt access tokens.
type JWTTokenVerifier struct {
	PublicKey *rsa.PublicKey
}

// Verify verifies a token and returns account id.
func (v *JWTTokenVerifier) Verify(token string) (string, error) {
	t, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{},
		func(*jwt.Token) (interface{}, error) {
			return v.PublicKey, nil
		})

	if err != nil {
		return "", fmt.Errorf("cannot parse token: %v", err)
	}

	if !t.Valid {
		return "", fmt.Errorf("token not valid")
	}

	clm, ok := t.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return "", fmt.Errorf("token claim is not RegisteredClaims")
	}

	if err := clm.Valid(); err != nil {
		return "", fmt.Errorf("claim not valid: %v", err)
	}

	return clm.Subject, nil
}

func (v *JWTTokenVerifier) TokenExpiresAt(token string) (int64, error) {
	t, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{},
		func(*jwt.Token) (interface{}, error) {
			return v.PublicKey, nil
		})

	if err != nil {
		return 0, fmt.Errorf("cannot parse token: %v", err)
	}

	clm, ok := t.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return 0, fmt.Errorf("token claim is not RegisteredClaims")
	}

	return clm.ExpiresAt.Unix(), nil
}
