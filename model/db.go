package model

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

func InitDB() {
	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	var err error
	db, err = sql.Open("postgres", "postgres://" + user + ":" + pass + "@localhost/" + dbname + "?sslmode=disable")

	if err != nil {
		fmt.Println("Error connecting to database", err)
	}
	err = db.Ping()

	if err != nil {
		fmt.Println("Could not ping database", err)
	}

}