package api

import (
	"net/http"
	"tmpl-go-vercel/app"
)

func Handle(w http.ResponseWriter, r *http.Request) { app.Handle(w, r) }
