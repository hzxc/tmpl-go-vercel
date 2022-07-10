package app

import (
	"net/http"
	"tmpl-go-vercel/app/grpc"
)

var handler http.HandlerFunc

func init() {
	grpcSrv := grpc.New()
	handler = func(w http.ResponseWriter, r *http.Request) {
		grpcSrv.ServeHTTP(w, r)
	}
}

func Handle(w http.ResponseWriter, r *http.Request) {
	handler.ServeHTTP(w, r)
}
