package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Welcome to Recipe Book!")
}

func RecipesIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	recipes := List()
	for _, recipe := range recipes {
		fmt.Fprintln(w, "Recipe: ", recipe.Name, recipe.Id)
	}
}

func RecipeGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
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
