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

// TODO: allow users to rate recipes on scale from 1-5. take the sum
//	and average of the ratings and assign to the recipe
