package jwt

import (
	"crypto/rsa"
	"encoding/base64"

	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
)

func ConvPrivKey(key string) *rsa.PrivateKey {
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

func ConvPubKey(key string) *rsa.PublicKey {
	keyBytes, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		zap.L().Fatal("cannot decode public key", zap.Error(err))
	}

	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(keyBytes)
	if err != nil {
		zap.L().Fatal("cannot parse public key", zap.Error(err))
	}

	return pubKey
}
