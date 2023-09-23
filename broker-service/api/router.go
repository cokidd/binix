package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func routes() http.Handler {
	r := mux.NewRouter()
	r.Use(corsMiddleware())
	r.HandleFunc("/", Broker).Methods(http.MethodGet)
	return r
}
