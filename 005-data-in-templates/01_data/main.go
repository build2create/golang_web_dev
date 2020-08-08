package main

import (
	"html/template"
	"os"
)

func main() {
	tpl := template.Must(template.ParseFiles("tmpl.gohtml"))
	tpl.ExecuteTemplate(os.Stdout, "tmpl.gohtml", 42)
}
