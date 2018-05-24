package main

type Recipe struct {
	Name        string   `json:"name"`
	Ingredients []string `json:"ingredients"`
}

type Recipes []Recipe
