package data

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var times int

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func connDB() *sql.DB {
	dsn := os.Getenv("DSN")
	for {
		conn, err := openDB(dsn)
		if err != nil {
			log.Println("Mysql not ready ...")
			times++
		} else {
			log.Println("Connected to Mysql")
			return conn
		}
		time.Sleep(6 * time.Second)
		if times > 5 {
			log.Println(err)
			return nil
		}
	}
}
