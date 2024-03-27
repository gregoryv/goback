package main

import (
	"embed"
	"html/template"
	"log"
	"os"
)

func main() {
	model := map[string]any{
		"Title": "<script>alert('WOW');</script>",
	}
	err := tpl.ExecuteTemplate(os.Stdout, "index.html", model)
	if err != nil {
		log.Fatal(err)
	}
}

var tpl *template.Template = template.Must(
	template.ParseFS(assets, "assets/*.html"),
)

//go:embed assets/*.html
var assets embed.FS
