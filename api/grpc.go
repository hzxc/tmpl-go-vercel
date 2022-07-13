package api

import (
	"fmt"
	"net/http"
	"tmpl-go-vercel/app"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Host:%s\n", r.Host)
	fmt.Printf("Method:%s\n", r.Method)
	fmt.Printf("RemoteAddr:%s\n", r.RemoteAddr)
	fmt.Printf("RequestURI:%s\n", r.RequestURI)
	app.Handle(w, r)
}
