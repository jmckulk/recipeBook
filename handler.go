package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/boltdb/bolt"
	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Recipe Book!")
}

func RecipesIndex(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// w.WriteHeader(http.StatusOK)
	// var recipes []*Recipe
	db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket([]byte("book")).Cursor()
		if c == nil {
			fmt.Fprintln(w, "add recipe using curl")
			return nil
		}
		for k, v := c.First(); k != nil; k, v = c.Next() {
			recipe, _ := decode(v)
			// recipes = append(recipes, recipe)
			fmt.Fprintln(w, "Recipe:", recipe.Name)
		}
		// if err := json.NewEncoder(w).Encode(recipes); err != nil {
		// 	panic(err)
		// }
		return nil
	})
}

func RecipeGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["recipeId"]
	recipe, _ := GetRecipe(id)
	fmt.Fprintln(w, "Recipe get:", recipe)
}

func RecipeCreate(w http.ResponseWriter, r *http.Request) {
	var recipe *Recipe
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Println("Recipe Create: Read from body")
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		log.Println("Recipe Create: close body")
		panic(err)
	}
	recipe, err = decode(body)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Println("Recipe Create: encode err")
			panic(err)
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, recipe)
	recipe.AddRecipe()
}
