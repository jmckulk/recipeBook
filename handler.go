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
	welcome := "Welcome to Recipe Book!"

	indexT.Execute(w, welcome)
}

func RecipesIndex(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// w.WriteHeader(http.StatusOK)
	recipes := List()
	recipeBook := struct {
		Book string
		List Recipes
	}{
		Book: "Recipe List",
		List: recipes,
	}
	listT.Execute(w, recipeBook)
	// if recipes == nil {
	// 	fmt.Fprintln(w, "Recipe book is empty.")
	// } else {
	// 	for _, recipe := range recipes {
	// 		fmt.Fprintln(w, "Recipe: ", recipe.Name, "\tCook Time: ", recipe.CookTime, "\tIngredients: ", recipe.IngredientList)
	// 	}
	// }
}

func RecipeCreate(w http.ResponseWriter, r *http.Request) {
	var recipe *Recipe
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	check(err)
	if err := r.Body.Close(); err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(body, &recipe)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatal(err)
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, recipe)
	recipe.AddRecipe()
}

func RecipeDelete(w http.ResponseWriter, r *http.Request) {
	var recipe *Recipe
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	check(err)
	if err := r.Body.Close(); err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(body, &recipe)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatal(err)
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, recipe)
	if err := DeleteRecipe(recipe.Name); err != nil {
		log.Fatal(err)
	}
}

func UpdateRecipeTime(w http.ResponseWriter, r *http.Request) {
	var recipe *Recipe
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	check(err)
	if err := r.Body.Close(); err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(body, &recipe)
	if recipe.CookTime == "" || recipe.Name == "" {
		log.Println("Need a time and recipe name to update recipe.")
	} else {
		if err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(422) // unprocessable entity
			if err := json.NewEncoder(w).Encode(err); err != nil {
				log.Fatal(err)
			}
		}
		err = UpdateTime(recipe.Name, recipe.CookTime)
		check(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func UpdateIngredients(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["recipeId"]
	ingredient := Ingredient(vars["ingredient"])
	UpdateIngredientList(id, ingredient)
	// Index(w, r)
}
