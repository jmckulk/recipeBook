package main

type Recipe struct {
	Name           string       `json:"name"`
	CookTime       string       `json:"cooktime"`
	Rating         int          `json:"rating"`
	IngredientList []Ingredient `json:"ingredientlist"`
}

type Ingredient struct {
	Name   string
	Amount string
}

type Recipes []Recipe
