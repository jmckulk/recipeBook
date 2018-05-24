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
		"POST",
		"/recipes",
		RecipeCreate,
	},
	Route{
		"RecipesDelete",
		"POST",
		"/recipes",
		RecipeDelete,
	},
}
