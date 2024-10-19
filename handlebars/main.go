package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	context := map[string]any{
		"Title": "Go Example Title",
	}
	template := template.Must(template.ParseFiles("index.html"))
	f, err := os.Create("output.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	err = template.Execute(f, context)
	if err != nil {
		log.Fatal(err)
	}
}
