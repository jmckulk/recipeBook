package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

// var db *bolt.DB
var t *template.Template
var check func(error)

const (
	DEFAULT_DB = "bolt"
)

func main() {
	var err error
	check = func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	err = InitTemplates()
	check(err)
	CreateDB(getDB())
	// r := &Recipe{Name: "Pork", CookTime: "30 min"}
	// r.AddRecipe()
	// UpdateTime("Pork", "60 min")
	// UpdateIngredientList("Pork", "Sauce")
	// UpdateIngredientList("Pork", "Meat")

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))

	CloseDB(getDB())
}

func CreateDB(db string) {
	switch db {
	case "bolt":
		log.Println("bolt")
	default:
		log.Println("pg")
	}
}

func CloseDB(db string) {
	switch db {
	case "bolt":
		// Close()
		log.Println("bolt")
	default:
		log.Println("pg")
	}
}
func getDB() string {
	var db string
	if db = os.Getenv("DATABASE"); len(db) == 0 {
		db = DEFAULT_DB
	}
	return db
}
