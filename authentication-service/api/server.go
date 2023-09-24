package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/cokidd/binix/authentication/data"
)

const webPort = "80"

type Auth struct {
	DB     *sql.DB
	Models data.Models
}

func Start() {
	log.Printf("Start broker service on port %s", webPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: Auth.routes(),
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
