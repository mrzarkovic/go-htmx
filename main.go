package main

import (
	"fmt"
	"fullstack/model"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
  
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}

	model.InitDB()

	fmt.Println("Starting server on port 8080")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		todos, err := model.GetTodos()
		if (err != nil) {
			fmt.Fprintf(w, "Error getting todos: %v\n", err)
		}

		fmt.Fprintf(w, "%v\n", todos)
	})

	http.ListenAndServe(":8080", nil)
}