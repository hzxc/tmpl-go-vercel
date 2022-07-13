package jwt

import (
	"crypto/rsa"
	"encoding/base64"

	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
)

func ConvKey(key string) *rsa.PrivateKey {
	keyBytes, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		zap.L().Fatal("cannot decode private key", zap.Error(err))
	}

	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(keyBytes)
	if err != nil {
		zap.L().Fatal("cannot parse private key", zap.Error(err))
	}

	return privKey
}
