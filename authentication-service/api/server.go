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
	log.Printf("Start Auth service on port %s", webPort)

	conn := data.ConnDB()
	if conn == nil {
		log.Panic("Can't connect to Mysql.")
	}

	auth := Auth{
		DB:     conn,
		Models: data.New(conn),
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: auth.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
