package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

func GetConnection() *sql.DB {
	db, err = sql.Open("postgres", "host=localhost port=5432 user=postgres password=suksesmulia07 dbname=orders_by sslmode=disable")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
