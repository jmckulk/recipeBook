package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// TODO: Setup handler functions to only perform api calls.
//	Functions will take ResponseWriter and Request and send the
//	information to the appropriate database.
// TODO: Add functionality to choose which database should be used
func Index(w http.ResponseWriter, r *http.Request) {
	message := "Welcome to Recipe Book!"
	indexT.Execute(w, message)
}

func RecipesIndex(w http.ResponseWriter, r *http.Request) {
	recipes := List()
	recipeBook := struct {
		Book string
		List Recipes
	}{
		Book: "Recipe List",
		List: recipes,
	}
	listT.Execute(w, recipeBook)
}

func RecipeCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		recipe := Recipe{
			Name:     r.FormValue("name"),
			CookTime: r.FormValue("cooktime"),
		}
		recipe.AddRecipe()
		data := r.FormValue("ingredients")
		ingredients := strings.Split(data, ",")
		for _, ingredient := range ingredients {
			newIngredient := Ingredient{
				Name:   ingredient,
				Amount: "0 Cups",
			}
			UpdateIngredientList(recipe.Name, newIngredient)
		}
		http.Redirect(w, r, "/recipes", http.StatusFound)
	} else {
		vars := mux.Vars(r)
		id := vars["recipeId"]
		if id != "" {
			recipe, err := GetRecipe(id)
			if err != nil {
				panic(err)
			}
			createT.Execute(w, recipe)
		} else {
			createT.Execute(w, nil)
		}
	}
}

func RecipeShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["recipeId"]
	recipe, err := GetRecipe(id)
	check(err)
	recipeT.Execute(w, recipe)
}

func RecipeDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["recipeId"]
	if err := DeleteRecipe(id); err != nil {
		log.Fatal(err)
	}
	http.Redirect(w, r, "/recipes", http.StatusFound)
	// var recipe *Recipe
	// body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	// check(err)
	// if err := r.Body.Close(); err != nil {
	// 	log.Fatal(err)
	// }
	// err = json.Unmarshal(body, &recipe)
	// if err != nil {
	// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// 	w.WriteHeader(422) // unprocessable entity
	// 	if err := json.NewEncoder(w).Encode(err); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// w.WriteHeader(http.StatusOK)
	// fmt.Fprintln(w, recipe)
	// if err := DeleteRecipe(recipe.Name); err != nil {
	// 	log.Fatal(err)
	// }
}

// func UpdateRecipeTime(w http.ResponseWriter, r *http.Request) {
// 	var recipe *Recipe
// 	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
// 	check(err)
// 	if err := r.Body.Close(); err != nil {
// 		log.Fatal(err)
// 	}
// 	err = json.Unmarshal(body, &recipe)
// 	if recipe.CookTime == "" || recipe.Name == "" {
// 		log.Println("Need a time and recipe name to update recipe.")
// 	} else {
// 		if err != nil {
// 			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 			w.WriteHeader(422) // unprocessable entity
// 			if err := json.NewEncoder(w).Encode(err); err != nil {
// 				log.Fatal(err)
// 			}
// 		}
// 		err = UpdateTime(recipe.Name, recipe.CookTime)
// 		check(err)
// 	}
// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 	w.WriteHeader(http.StatusOK)
// }

// func UpdateIngredients(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id := vars["recipeId"]
// 	ingredient := Ingredient{
// 		Name:   vars["ingredient"],
// 		Amount: "0 Cups",
// 	}
// 	UpdateIngredientList(id, ingredient)
// 	http.Redirect(w, r, "/recipes", http.StatusFound)
// }
