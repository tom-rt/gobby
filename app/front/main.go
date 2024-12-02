package main

import (
	"fmt"
	"net/http"

	templlib "github.com/a-h/templ"
	"github.com/tom-rt/gobby/templ"
)

func main() {
	component := templ.Hello("toto")
	http.Handle("/", templlib.Handler(component))

	fmt.Println("listening on port 3000")
	http.ListenAndServe(":3000", nil)
}
