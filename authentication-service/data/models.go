package data

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Models struct {
	User       User
	Role       Role
	Permission Permission
}

type User struct {
}

type Role struct {
}

type Permission struct {
}

func New(dbPool *sql.DB) Models {
	db = dbPool
	return Models{
		User:       User{},
		Role:       Role{},
		Permission: Permission{},
	}
}
