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

const list = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="UTF-8">
    <title>{{.Book}}</title>
  </head>
  <body>
    <h1>Recipe List</h1>
    {{range .List}}
      <dl>
        <dt>{{.Name}}</dt>
        <dd>-{{.CookTime}}</dd>
        
      </dl>




      {{else}}<div><strong>No Recipes</strong></div>{{end}}
  </body>
</html>
`

var indexT *template.Template
var listT *template.Template

func InitTemplates() error {
	var err error
	indexT, err = template.New("Index").Parse(index)
	if err != nil {
		return err
	}
	listT, err = template.New("List").Parse(list)
	if err != nil {
		return nil
	}
	return err
}
