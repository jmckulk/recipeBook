package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

// TODO: Binary build pack
// var db *bolt.DB
var t *template.Template
var check func(error)

const (
	DEFAULT_DB   = "pg"
	DEFAULT_PORT = "8080"
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
	OpenDB(getDB())
	// recipe := Recipe{
	// 	Name:     "Pork",
	// 	CookTime: "30 min",
	// 	Rating:   5,
	// }
	// recipe.AddRecipe()
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":"+getPort(), router))

	CloseDB(getDB())
}

func OpenDB(db string) {
	switch db {
	case "bolt":
		OpenBolt()
	case "pg":
		OpenPG()
	}
}

func CloseDB(db string) {
	switch db {
	case "bolt":
		CloseBolt()
	case "pg":
		ClosePG()
	default:
		log.Fatal("Invalid DB")
	}
}

func (r *Recipe) AddRecipe() {
	switch getDB() {
	case "bolt":
		r.AddRecipeBolt()
	case "pg":
		r.AddRecipePG()
	}
}

func GetRecipe(id string) (*Recipe, error) {
	switch getDB() {
	case "bolt":
		return GetRecipeBolt(id)
	case "pg":
		return GetRecipePG(id)
	default:
		return nil, nil
	}
}

func List() []Recipe {
	switch getDB() {
	case "bolt":
		return ListBolt()
	case "pg":
		return ListPG()
	default:
		return nil
	}
}

func DeleteRecipe(id string) error {
	switch getDB() {
	case "bolt":
		return DeleteRecipeBolt(id)
	case "pg":
		return DeleteRecipePG(id)
	default:
		return nil
	}
}

func getDB() string {
	var db string
	if db = os.Getenv("DATABASE"); len(db) == 0 {
		db = DEFAULT_DB
	}
	return db
}

func getPort() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = DEFAULT_PORT
	}
	return port
}
