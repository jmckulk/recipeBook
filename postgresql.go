package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var pgdb *sql.DB

const (
	db_name = "recipe_book"
)

func OpenPG() {
	var err error
	// if pgdb == nil {
	// 	log.Fatal("no database")
	// }
	dbinfo := fmt.Sprintf("dbname=%s sslmode=%s", db_name, "disable")
	pgdb, err = sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}
	err = pgdb.Ping()
	if err != nil {
		log.Fatal(err)
	}
	createPGDB()
}

func ClosePG() {
	pgdb.Close()
}

func createPGDB() {
	if _, err := pgdb.Exec("CREATE TABLE IF NOT EXISTS recipes (name text, data json)"); err != nil {
		log.Fatal(err)
	}
}

func (r *Recipe) AddRecipePG() error {
	recipe, err := json.Marshal(r)
	str := fmt.Sprintf("INSERT INTO recipes (\"name\", \"data\") VALUES ('%s', '%s')", r.Name, recipe)
	if _, err := pgdb.Exec(str); err != nil {
		log.Fatal(err)
	}
	return err
}

func GetRecipePG(id string) (*Recipe, error) {
	var r []byte
	str := fmt.Sprintf("select data from recipes where name = '%s'", id)
	err := pgdb.QueryRow(str).Scan(&r)
	if err != nil {
		log.Fatal("Get: ", err)
	}
	var recipe *Recipe
	err = json.Unmarshal(r, &recipe)
	if err != nil {
		log.Fatal(err)
	}
	return recipe, nil
}

func ListPG() []Recipe {
	var recipes []Recipe
	var data []byte
	str := fmt.Sprintf("select data from recipes")
	r, err := pgdb.Query(str)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()
	for r.Next() {
		err := r.Scan(&data)
		if err != nil {
			log.Fatal(err)
		}
		var recipe Recipe
		err = json.Unmarshal(data, &recipe)
		if err != nil {
			log.Fatal(err)
		}
		recipes = append(recipes, recipe)
	}
	return recipes
}

func DeleteRecipePG(id string) error {
	str := fmt.Sprintf("delete from recipes where name='%s'", id)
	if _, err := pgdb.Exec(str); err != nil {
		log.Fatal(err)
	}
	return nil
}
