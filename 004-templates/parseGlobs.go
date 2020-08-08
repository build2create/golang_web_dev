package main

import (
	"log"
	"os"
	"text/template"
)

func parseGlobs() {
	tpl := template.Must(template.ParseFiles("templateSamples/tone.txt", "templateSamples/ttwo.txt", "templateSamples/tthree.txt"))

	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal(err)
	}

	tpl.ExecuteTemplate(os.Stdout, "tone.txt", nil)
	if err != nil {
		log.Fatal(err)
	}
	tpl.ExecuteTemplate(os.Stdout, "ttwo.txt", nil)
	if err != nil {
		log.Fatal(err)
	}
	tpl.ExecuteTemplate(os.Stdout, "tthree.txt", nil)
	if err != nil {
		log.Fatal(err)
	}
}
