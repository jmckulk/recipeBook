package main

type Recipe struct {
	Name     string `json:"name"`
	CookTime string `json:"cooktime"`
}

type Recipes []Recipe
