package cors

import (
	"net/http"
	"net/url"
	"strings"

	_env "tmpl-go-vercel/app/utils/env"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
)

type Handler struct {
	options *options
	reg     PatternRegistry
}

func NewHandler(r PatternRegistry, opts ...Option) *Handler {
	return &Handler{reg: r, options: evaluateOptions(opts)}
}

func (r *Handler) RegisterHandler(mux *runtime.ServeMux) {
	for _, p := range r.reg {
		mux.Handle("OPTIONS", p, r.ServeHTTP)
	}
}

func (r Handler) ServeHTTP(w http.ResponseWriter, req *http.Request, _ map[string]string) {
	headers := map[string]string{
		"access-control-allow-methods": "POST,GET,OPTIONS,PUT,DELETE",
		"access-control-allow-headers": req.Header.Get("access-control-request-headers"),
	}
	if r.options.allowHost == "*" {
		headers["access-control-allow-origin"] = "*"
	} else if r.options.allowHost != "" {
		origin := req.Header.Get("Origin")

		u, err := url.Parse(origin)
		if err != nil {
			zap.S().Errorw("Failed to parse CORS origin header", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		ok := u.Host == r.options.allowHost ||
			(r.options.allowSubdomain && strings.HasSuffix(u.Host, "."+r.options.allowHost))
		if !ok {

			zap.S().Errorf("CORS request from prohibited domain %v", origin)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !_env.FromHost().DevMode() {
			u.Scheme = "https"
		}
		headers["access-control-allow-origin"] = u.String()
		headers["access-control-allow-credentials"] = "true"
		headers["vary"] = "Origin"
	}
	for k, v := range headers {
		w.Header().Set(k, v)
	}
}
