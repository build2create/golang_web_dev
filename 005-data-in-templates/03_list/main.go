package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	tpl := template.Must(template.ParseFiles("WorldLeaders.html"))

	err := tpl.ExecuteTemplate(os.Stdout, "WorldLeaders.html", worldLeaders)
	if err != nil {
		log.Fatalln(err)
	}
}
