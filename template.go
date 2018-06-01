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
	// indexT = template.Must(template.New("index").Parse(indexHTML))
	// listT = template.Must(template.New("list").Parse(listHTML))
	// createT = template.Must(template.New("create").Parse(createHTML))
	// recipeT = template.Must(template.New("recipe").Parse(recipeHTML))
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

// var indexHTML = `
// <body>
//   <h1>{{.}}</h1>
//   <a href="/recipes">Recipes List</a></br>
//   <a href="/recipes/create">Create Recipe</a>
// </body>
// `
//
// var listHTML = `
// <head>
// <style>
// table, th, td {
//   border: 1px solid #dddddd;
// }
// th, td {
//   padding: 5px;
// }
// td {
//   text-align: center;
// }
// tr:nth-child(even) {
//   background-color: #dddddd;
// }
// </style>
// <meta charset="UTF-8">
// <title>{{.Book}}</title>
// </head>
// <body>
//   <table>
//     <caption><b>Recipe List</b></caption>
//     <tr>
//       <th>Name</th>
//       <th>Cook Time</th>
//       <th>Rating</th>
//     </tr>
//     {{range .List}}
//       <tr>
//         <td><a href="/recipes/{{.Name}}">{{.Name}}</td>
//         <td>{{.CookTime}}</td>
//         <td>{{.Rating}}</td>
//       </tr>
//     {{end}}
//   </table>
//   <form action="/">
//     <input type="submit" value="Home"/>
//   </form>
// </body>
// `
// var createHTML = `
// <body>
//   <form action="/">
//     <input type="submit" value="Home"/>
//   </form>
//   <form action="/recipes/create" method="POST">
//     Recipe name:<br>
//     <input type="text" name="name" value={{.Name}}><br>
//     Cook Time:<br>
//     <input type="text" name="cooktime" value={{.CookTime}}><br>
//     Ingredients:<br>
//     <input type="text" name="ingredients"><br>
//     <input type="submit" value="Submit">
//   </form>
// </body>
// `
//
// var recipeHTML = `
// <head>
// <style>
// table, th, td {
//   border: 1px solid #dddddd;
// }
// th, td {
//   padding: 5px;
// }
// td {
//   text-align: center;
// }
// tr:nth-child(even) {
//   background-color: #dddddd;
// }
// </style>
// </head>
// <body>
//   <h1>{{.Name}}</h1>
//   <div>Cook Time: {{.CookTime}}</div>
//   <div>Rating: {{.Rating}}</div>
//   </br>
//   <table>
//     <caption><b>Ingredients:</b></caption>
//     <tr>
//       <th>Name</th>
//       <th>Amount</th>
//     </tr>
//     {{range .IngredientList}}
//     <tr>
//       <td>{{.Name}}</td>
//       <td>{{.Amount}}</td>
//     </tr>
//     {{end}}
//   </table>
//   </br>
//   <form action="/recipes/delete/{{.Name}}" method="POST">
//     <input type="submit" value="Delete"/>
//   </form>
//   <form action="/recipes/create/{{.Name}}" method="GET">
//     <input type="submit" value="Update"/>
//   </form>
//   <form action="/" method="GET">
//     <input type="submit" value="HOME"/>
//   </form>
// </body>
// `
