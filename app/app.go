package app

import (
	"net/http"
	"tmpl-go-vercel/app/global"
	"tmpl-go-vercel/app/grpc"
	_ "tmpl-go-vercel/app/init"
	_ "tmpl-go-vercel/app/services"

	"go.uber.org/zap"
)

var (
	// handler http.HandlerFunc
	grpcSrv grpc.Handler
)

func init() {
	var (
		err error
	)

	o := grpc.NewServerOptions(false, global.ZapLogger)

	grpcSrv, err = o.New()
	if err != nil {
		zap.L().Fatal(err.Error())
	}
}

func Handle(w http.ResponseWriter, r *http.Request) {
	grpcSrv.ServeHTTP(w, r)
}
