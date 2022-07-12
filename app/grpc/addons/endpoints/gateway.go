package endpoints

import (
	"context"
	"reflect"
	gort "runtime"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type RegisterProxyHandlerFunc func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)

type proxyHandler struct {
	Register RegisterProxyHandlerFunc
}

type ProxyRegistry []*proxyHandler

func (r *ProxyRegistry) Register(fn RegisterProxyHandlerFunc) {
	if *r == nil {
		*r = make([]*proxyHandler, 0)
	}
	*r = append(*r, &proxyHandler{
		Register: fn,
	})
}

func (r ProxyRegistry) ApplyTo(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	for _, ep := range r {
		zap.S().Infof("Registering grpc-gateway endpoint: %s", funcName(ep.Register))
		if err := ep.Register(context.Background(), mux, endpoint, opts); err != nil {
			return nil
		}
	}
	return nil
}

func funcName(i interface{}) string {
	name := gort.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	return name[strings.LastIndex(name, ".")+1:]
}
