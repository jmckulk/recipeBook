package main

type Recipe struct {
	Name           string       `json:"name"`
	CookTime       string       `json:"cooktime"`
	IngredientList []Ingredient `json:"ingredientlist"`
}

type Ingredient string
type Recipes []Recipe
