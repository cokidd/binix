package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (a Auth) routes() http.Handler {
	r := mux.NewRouter()
	r.Use(corsMiddleware())

	return r
}
