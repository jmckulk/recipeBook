package main

import "html/template"

var indexT *template.Template
var listT *template.Template
var createT *template.Template
var recipeT *template.Template

func InitTemplates() error {
	var err error
	indexT = template.Must(template.ParseFiles("HTML/Index.html"))
	listT = template.Must(template.ParseFiles("HTML/List.html"))
	createT = template.Must(template.ParseFiles("HTML/Create.html"))
	recipeT = template.Must(template.ParseFiles("HTML/Recipe.html"))
	return err
}

// TODO: Add functions to format information from form and format it to
//  json. Then pass the json to the appropriate api function. This will
//  remove front end handling from the api
// TODO: Update recipe html to have a button leading to create recipe form.
// TODO: Update create recipe functions to go to the specific recipe page
//  after creating or updating recipe
// TODO: Add functionality to pull data into the create form when page is
//  pulled up for a specific recipe
// TODO: Add functionality to have recipe instructions
// TODO: Add home button to pages
