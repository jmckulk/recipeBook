package main

import "html/template"

var indexT *template.Template
var listT *template.Template
var createT *template.Template

func InitTemplates() error {
	var err error
	indexT = template.Must(template.ParseFiles("HTML/Index.html"))
	listT = template.Must(template.ParseFiles("HTML/List.html"))
	createT = template.Must(template.ParseFiles("HTML/Create.html"))
	return err
}
