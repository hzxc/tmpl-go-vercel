package app

import (
	"fmt"
	"net/http"
	"tmpl-go-vercel/app/grpc"

	"go.uber.org/zap"
)

var (
	// handler http.HandlerFunc
	grpcSrv grpc.Handler
)

func init() {

	zapLogger, err := zap.NewDevelopment(zap.AddCaller())
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	zap.ReplaceGlobals(zapLogger)

	zap.S().Debug("init grpc")
	zap.S().Info("init grpc")
	o := grpc.NewServerOptions(false, zapLogger)

	grpcSrv, err = o.New()
	if err != nil {
		zap.L().Fatal(err.Error())
	}

	// handler = func(w http.ResponseWriter, r *http.Request) {
	// 	zap.L().Info(r.URL.Port())
	// 	zap.L().Info(r.Host)
	// 	zap.L().Info(r.RemoteAddr)
	// 	zap.L().Info(r.RequestURI)
	// 	grpcSrv.ServeHTTP(w, r)
	// }
}

func Handle(w http.ResponseWriter, r *http.Request) {
	zap.L().Info(r.URL.Port())
	zap.L().Info(r.Host)
	zap.L().Info(r.RemoteAddr)
	zap.L().Info(r.RequestURI)
	// hd, _ := json.Marshal(r.Header)
	// zap.L().Info(string(hd))
	// handler.ServeHTTP(w, r)
	grpcSrv.ServeHTTP(w, r)
}
