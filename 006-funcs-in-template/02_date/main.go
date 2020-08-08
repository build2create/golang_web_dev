package main

import (
	"fmt"
	"html/template"
	"os"
	"time"
)

var fm = template.FuncMap{
	"ft": formatTime,
}

func formatTime(t time.Time) string {
	return fmt.Sprintf("Day: %d, Month: %s, Year: %d", t.Day(), t.Month().String(), t.Year())
}
func main() {
	tpl := template.Must(template.New("").Funcs(fm).ParseFiles("tmpl.gohtml"))
	f, _ := os.Create("time.html")
	tpl.ExecuteTemplate(f, "tmpl.gohtml", time.Now())
}
