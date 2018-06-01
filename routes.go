package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"/",
		Index,
	},
	Route{
		"Recipes",
		"/recipes",
		RecipesIndex,
	},
	Route{
		"RecipesCreate",
		"/recipes/create/{recipeId}",
		RecipeCreate,
	},
	Route{
		"RecipeCreate",
		"/recipes/create",
		RecipeCreate,
	},
	Route{
		"Recipes",
		"/recipes/{recipeId}",
		RecipeShow,
	},
	Route{
		"RecipesDelete",
		"/recipes/delete/{recipeId}",
		RecipeDelete,
	},
}
