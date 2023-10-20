package model

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

func InitDB() {
	var err error
	db, err = sql.Open("postgres", "postgres://todo:todo@localhost/todo?sslmode=disable")

	if err != nil {
		fmt.Println("Error connecting to database", err)
	}
	err = db.Ping()

	if err != nil {
		fmt.Println("Could not ping database", err)
	}

}