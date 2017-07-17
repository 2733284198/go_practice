package main

// https://gowebexamples.github.io/templates/
// http://studygolang.com/articles/8277
// http://studygolang.com/articles/3071

import (
	"html/template"
	"net/http"
)

type Todo struct {
	Task string
	Done bool
}

func main() {
	tmpl := template.Must(template.ParseFiles("./tmpl/todos.html"))
	todos := []Todo{
		{"Learn Go", true},
		{"Read Go Web Examples", true},
		{"Create a web app in Go", false},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, struct{ Todos []Todo }{todos})
	})

	http.ListenAndServe(":8080", nil)
}
