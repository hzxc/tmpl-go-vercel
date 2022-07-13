package auth

import "crypto/rsa"

type options struct {
	pubKey *rsa.PublicKey
}

type Option func(*options)

func PubKey(pk *rsa.PublicKey) Option {
	return func(o *options) {
		o.pubKey = pk
	}
}

func evaluateOptions(opts []Option) *options {
	optCopy := &options{}
	for _, o := range opts {
		o(optCopy)
	}
	return optCopy
}
