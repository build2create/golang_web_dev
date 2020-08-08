package main

import (
	"html/template"
	"os"
)

func main() {
	tpl := template.Must(template.ParseFiles("tpl.gohtml"))
	tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", `Well this is a variable in go template`)
}
