package api

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

func Start() {
	log.Printf("Start itam service on port %s", webPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: routes(),
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
