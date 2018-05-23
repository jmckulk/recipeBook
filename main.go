package main

import (
	"log"
	"net/http"
)

// var db *bolt.DB

func main() {
	Open()
	// r := &Recipe{Name: "Pork", Id: "0"}
	// r.AddRecipe()
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))

	Close()
}
