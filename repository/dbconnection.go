package repository

import (
	"database/sql"
)

func DatabaseConnect() *sql.DB {
	connection := "user=postgres dbname=products password=password host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}
	return db
}
