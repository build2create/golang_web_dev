package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	tpl.ExecuteTemplate(w, "template.gohtml", req.Form)
}
func init() {
	tpl = template.Must(template.ParseFiles("template.gohtml"))
}
func main() {
	var m hotdog
	http.ListenAndServe(":8080", m)
}
