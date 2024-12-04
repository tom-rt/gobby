package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/tom-rt/gobby/template"
)

func main() {
	component := template.Hello("toto")
	http.Handle("/", templ.Handler(component))

	// server static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", fs)

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
