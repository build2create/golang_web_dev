package main

import (
	"html/template"
	"os"
	"strings"
)

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s
}
func main() {
	f := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}

	c := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}
	cars := []car{f, c}
	tpl := template.Must(template.New("").Funcs(fm).ParseFiles("tmpl.gohtml"))
	tpl.ExecuteTemplate(os.Stdout, "tmpl.gohtml", cars)
}
