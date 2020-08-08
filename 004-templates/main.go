package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	//simpleTemplate()
	// Part 1: html template to os.Stdout
	tpl, err := template.ParseFiles("tplgo.html")
	if err != nil {
		log.Fatal("Error in parsing")
	}
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal("Error in executing")
	}
	//Part 2 : html template to new file index.html
	ns, err := os.Create("index.html")
	if err != nil {

	}
	defer ns.Close()
	err = tpl.Execute(ns, nil)
	if err != nil {
		log.Fatal("Error in executing 2")
	}
	//Part 3: create 3 files one, two and three and examine the template funcs
	examineTemplate()
	//Part4 : create 3 files in a folder and then try to access them using ParseGlob
	//parseGlobs()
}
