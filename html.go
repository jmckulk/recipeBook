package main

import "html/template"

const index = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="UTF-8">
    <title>{{.}}</title>
  </head>
  <body>
    <h1>{{.}}</h1>
  </body>
</html>`

var indexT *template.Template

func InitTemplates() error {
	var err error
	indexT, err = template.New("Index").Parse(index)

	return err
}
