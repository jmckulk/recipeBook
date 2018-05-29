package main

import (
	"html/template"
	"log"
	"net/http"
)

// var db *bolt.DB
var t *template.Template
var check func(error)

func main() {
	var err error
	check = func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}

	err = InitTemplates()
	check(err)
	Open()
	// r := &Recipe{Name: "Pork", CookTime: "30 min"}
	// r.AddRecipe()
	// UpdateTime("Pork", "60 min")
	// UpdateIngredientList("Pork", "Sauce")
	// UpdateIngredientList("Pork", "Meat")

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))

	Close()
}
