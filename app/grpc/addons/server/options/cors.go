/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package options

import (
	"tmpl-go-vercel/app/grpc/addons/server"
)

type CorsOptions struct {
	Enable         bool
	OriginHost     string
	AllowSubdomain bool
}

func NewCORSOptions() *CorsOptions {
	return &CorsOptions{
		// Enable CORS support
		Enable: true,
		// Allowed CORS origin host e.g, domain[:port]
		OriginHost: "*",
		// Allow CORS request from subdomains of origin
		AllowSubdomain: true,
	}
}

func (o *CorsOptions) ApplyTo(cfg *server.Config) error {
	cfg.EnableCORS = o.Enable
	cfg.CORSOriginHost = o.OriginHost
	cfg.CORSAllowSubdomain = o.AllowSubdomain

	return nil
}

func (o *CorsOptions) Validate() []error {
	return nil
}
