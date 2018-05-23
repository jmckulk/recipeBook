package main

type Recipe struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Recipes []Recipe
