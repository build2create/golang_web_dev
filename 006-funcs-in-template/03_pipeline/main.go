package main

import (
	"html/template"
	"math"
	"os"
)

func doubleNumber(x int) int {
	return x + x
}
func squareNumber(x int) int {
	return x * x
}
func squareRoot(x int) float64 {
	return math.Sqrt(float64(x))
}

var fm = template.FuncMap{
	"double":     doubleNumber,
	"square":     squareNumber,
	"squareRoot": squareRoot,
}

func main() {
	tpl := template.Must(template.New("").Funcs(fm).ParseFiles("math.gohtml"))
	tpl.ExecuteTemplate(os.Stdout, "math.gohtml", 42)
}
