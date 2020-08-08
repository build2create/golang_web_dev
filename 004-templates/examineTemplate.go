package main

import (
	"html/template"
	"log"
	"os"
)

func examineTemplate() {
	tpl, err := template.ParseFiles("one.txt")
	if err != nil {
		log.Fatalln("Error in parsing one.txt", err)
	}
	//Execute one.txt
	// nil passed as a second argument as there is no parsed data object to apply the template
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln("Error in executing one.txt", err)
	}
	//parse two files at once
	tpl, err = template.ParseFiles("two.txt", "three.txt")
	if err != nil {
		log.Fatalln("Error in parsing two.txt, three.txt", err)
	}
	// execute the template with file three.txt first
	err = tpl.ExecuteTemplate(os.Stdout, "three.txt", nil)
	if err != nil {
		log.Fatalln("Error in executing three.txt", err)
	}
	// execute the template with file two.txt next
	err = tpl.ExecuteTemplate(os.Stdout, "two.txt", nil)
	if err != nil {
		log.Fatalln("Error in executing two.txt", err)
	}
	// execute the template with file one.txt again
	//ExecuteTemplate exits silently
	/*err = tpl.ExecuteTemplate(os.Stdout, "one.txt", nil)
	if err != nil {
		log.Fatalln("Error in executing one.txt again", err)
	}
	*/
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln("Hello::", err)
	}
}
