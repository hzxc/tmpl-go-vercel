package app

import (
	"fmt"
	"net/http"
	"tmpl-go-vercel/app/grpc"

	"go.uber.org/zap"
)

var handler http.HandlerFunc

func init() {
	zapLogger, err := zap.NewProduction(zap.AddCaller())
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	zap.ReplaceGlobals(zapLogger)

	o := grpc.NewServerOptions(false, zapLogger)
	o.RecommendedOptions.Cors.AllowSubdomain = true
	o.RecommendedOptions.Cors.OriginHost = "*"
	o.RecommendedOptions.Cors.Enable = true

	grpcSrv, err := o.New()
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
