package main

import (
	"fmt"
	"net/http"
)

type vadapav int

func (m vadapav) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Any code goes here")
}
func main() {
	var d vadapav
	http.ListenAndServe(":8080", d)
}
