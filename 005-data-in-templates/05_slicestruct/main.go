package main

import (
	"html/template"
	"os"
)

type game struct {
	Name string
	Rank int
}

func main() {
	g1 := game{
		Name: "Football",
		Rank: 1,
	}
	g2 := game{
		Name: "Tennis",
		Rank: 2,
	}
	g3 := game{
		Name: "Cricket",
		Rank: 3,
	}
	g4 := game{
		Name: "Basketball",
		Rank: 4,
	}
	g5 := game{
		Name: "Hockey",
		Rank: 5,
	}
	games := []game{g1, g2, g3, g4, g5}
	tpl := template.Must(template.ParseFiles("Games.gohtml"))
	//Create a file to write a template
	f, err := os.Create("./Games.html")
	if err != nil {
		panic(err)
	}
	tpl.ExecuteTemplate(f, "Games.gohtml", games)

}
