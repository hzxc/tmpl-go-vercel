package options

import (
	"tmpl-go-vercel/app/grpc/addons/server"
)

type SecureServingOptions struct {
	SecureAddr    string
	PlaintextAddr string
	APIDomain     string
	CACertFile    string
	CertFile      string
	KeyFile       string
}

func NewSecureServingOptions() *SecureServingOptions {
	return &SecureServingOptions{
		// host:port used to serve secure apis
		SecureAddr: ":8443",
		// host:port used to serve http json apis
		PlaintextAddr: ":8080",
	}
}

func (o *SecureServingOptions) ApplyTo(cfg *server.Config) error {
	cfg.SecureAddr = o.SecureAddr
	cfg.PlaintextAddr = o.PlaintextAddr
	cfg.APIDomain = o.APIDomain
	cfg.CACertFile = o.CACertFile
	cfg.CertFile = o.CertFile
	cfg.KeyFile = o.KeyFile

	return nil
}

func (o *SecureServingOptions) Validate() []error {
	return nil
}
