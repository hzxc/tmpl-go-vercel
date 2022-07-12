package options

import (
	"tmpl-go-vercel/app/grpc/addons/server"
)

type RecommendedOptions struct {
	Cors          *CorsOptions
	SecureServing *SecureServingOptions
}

func NewRecommendedOptions() *RecommendedOptions {
	return &RecommendedOptions{
		Cors:          NewCORSOptions(),
		SecureServing: NewSecureServingOptions(),
	}
}

func (o *RecommendedOptions) ApplyTo(config *server.Config) error {
	if err := o.Cors.ApplyTo(config); err != nil {
		return err
	}
	if err := o.SecureServing.ApplyTo(config); err != nil {
		return err
	}
	return nil
}

func (o *RecommendedOptions) Validate() []error {
	var errors []error
	errors = append(errors, o.Cors.Validate()...)
	errors = append(errors, o.SecureServing.Validate()...)

	return errors
}
