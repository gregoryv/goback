package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	tpl, err := template.ParseFiles("examples/index.html")
	if err != nil {
		log.Fatal(err)
	}

	model := map[string]any{
		"Title": "<script>alert('WOW');</script>",
	}
	if err := tpl.Execute(os.Stdout, model); err != nil {
		log.Fatal(err)
	}
}
