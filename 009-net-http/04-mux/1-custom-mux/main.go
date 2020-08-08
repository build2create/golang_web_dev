package main

import (
	"fmt"
	"net/http"
)

type samosa struct{}
type icecream struct{}

func (s samosa) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Samosa")
}
func (i icecream) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Ice-Cream")
}
func main() {
	var s samosa
	var i icecream

	mux := http.NewServeMux()
	mux.Handle("/samosa", s)
	mux.Handle("/icecream", i)
	http.ListenAndServe(":8080", mux)
}
