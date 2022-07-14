package app

import (
	"net/http"
	"tmpl-go-vercel/app/global"
	"tmpl-go-vercel/app/grpc"

	_ "tmpl-go-vercel/app/init"
	_ "tmpl-go-vercel/app/services"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"go.uber.org/zap"
)

var (
	handler http.HandlerFunc
)

func init() {
	var (
		err     error
		grpcSrv *grpcweb.WrappedGrpcServer
	)

	o := grpc.NewServerOptions(true, global.ZapLogger)

	grpcSrv, err = o.New()
	if err != nil {
		zap.L().Fatal(err.Error())
	}

	handler = func(w http.ResponseWriter, r *http.Request) {
		grpcSrv.ServeHTTP(w, r)
	}
}

func Handle(w http.ResponseWriter, r *http.Request) {
	handler.ServeHTTP(w, r)
}
