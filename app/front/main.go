package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	component := hello("Templ!")
	http.Handle("/", templ.Handler(component))

	fmt.Println("Listening on port 3000 🚀")
	http.ListenAndServe(":3000", nil)
}
