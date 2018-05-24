package main

import (
	"log"
	"net/http"
)

// var db *bolt.DB

func main() {
	Open()
	// r := &Recipe{Name: "Pork", CookTime: "30 min"}
	// r.AddRecipe()
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))

	Close()
}
