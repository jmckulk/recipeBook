package main

type Recipe struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Recipes []Recipe
