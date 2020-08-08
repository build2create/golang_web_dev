package main

import (
	"fmt"
	"net/http"
)

func s(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Samosa")
}
func i(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Ice-Cream")
}
func main() {

	http.HandleFunc("/samosa", s)
	http.HandleFunc("/icecream", i)
	http.ListenAndServe(":8080", nil)
}
