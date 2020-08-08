package main

import (
	"log"
	"os"
)

func main() {
	name := "Sunday"
	tpl := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello World!</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`
	file, error := os.Create("FirstGo.html")
	if error != nil {
		log.Fatal("Failed opening the file")
	}
	_, error = file.Write([]byte(tpl))
	if error != nil {
		log.Fatal("Error writing to the file")
	}
	/*
		r := strings.NewReader(tpl)
		if _, err := io.Copy(file, r); err != nil {
			log.Fatal("Fail to copy the string")
		}
	*/
}
