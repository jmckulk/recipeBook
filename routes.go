package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"Recipes",
		"GET",
		"/recipes",
		RecipesIndex,
	},
	Route{
		"RecipesCreate",
		"GET",
		"/recipes/createform",
		RecipeCreateForm,
	},
	Route{
		"RecipesCreate",
		"POST",
		"/recipes/create",
		RecipeCreate,
	},
	Route{
		"Recipes",
		"GET",
		"/recipes/{recipeId}",
		RecipeShow,
	},
	Route{
		"RecipesDelete",
		"POST",
		"/recipes/delete/{recipeId}",
		RecipeDelete,
	},
	// Route{
	// 	"UpdateRecipeTime",
	// 	"POST",
	// 	"/recipes/update/time",
	// 	UpdateRecipeTime,
	// },
}
