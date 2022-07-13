package options

import (
	"tmpl-go-vercel/app/grpc/server"
)

type RecommendedOptions struct {
	Cors *CorsOptions
}

func NewRecommendedOptions() *RecommendedOptions {
	return &RecommendedOptions{
		Cors: NewCORSOptions(),
	}
}

func (o *RecommendedOptions) ApplyTo(config *server.Config) error {
	if err := o.Cors.ApplyTo(config); err != nil {
		return err
	}

	return nil
}

func (o *RecommendedOptions) Validate() []error {
	var errors []error
	errors = append(errors, o.Cors.Validate()...)

	return errors
}
