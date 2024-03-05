package main

import (
	"fmt"
	"strings"

	"github.com/gregoryv/deck"
	. "github.com/gregoryv/web"
)

func main() {
	p := deck.Deck{
		Title:     "Golang; Backend development",
		Author:    "Gregory Vinčić",
		AutoCover: true,
		AutoTOC:   true,
	}
	// dark mode
	// p.Style("html, body, a", "background-color: #2E2E34", "color: #f0f8ff")
	p.Style(".header", "background-color: #f6f5f4", "border-bottom: 1px inset #000")
	p.Style(".page .footer", "bottom: 1vh") // slightly up
	// center toc if short titles
	// p.Style(".toc", "margin-left: 10vw", "width: 20vw")
	p.Style(".shell",
		"padding: 1em",
		"border-radius: 10px",
		"min-width: 40vw",
		"overflow: wrap",
	)
	p.Style(".dark",
		"background-color: #2e2e34",
		"color: aliceblue",
	)
	p.Style("nav",
		"column-count: 2", // Fix this at the end with a manual toc
		"font-size: 0.8em",
	)
	p.Style(".srcfile ol",
		"padding-right: 40px",
	)
	p.Style(".shell",
		"font-size: 0.7em",
	)
	p.Style("p>a",
		"text-decoration: underline",
	)
	p.Style("quote",
		"display: inline-block",
		"font-style: italic",
		"padding-left: 2vw",
		"padding-right: 2vw",
		"width: 60vw",
	)
	p.Style("quote>a",
		"font-size: 0.7em",
		"float: right",
	)
	// ----------------------------------------  ----------------------------------------
	// ----------------------------------------  ----------------------------------------
	p.NewCard(
		H2("Quick start"),
		deck.Middle(70,
			P(`Download and install Go in the $HOME directory.`),
			Pre(Class("shell dark"),
				`$ cd $HOME
$ wget https://go.dev/dl/go1.22.0.linux-amd64.tar.gz
$ tar xvfz go1.22.0.linux-amd64.tar.gz
$ export PATH="$PATH:$HOME/go/bin"
$ go version
go version go1.22.0 linux/amd64`,
			),

			P("For more information refer to ",
				A(Href("https://go.dev/doc/install"),
					`"Download and install"`),
				".",
			),
		),
	)

	// ----------------------------------------
	p.NewCard(
		H2("Concepts"),
		deck.Middle(70,

			P(`In Go; <b>modules</b> are used to group
			<b>packages</b> and functions.`),

			quote(`You can collect related packages into modules, then
		publish the modules for other developers to use.`,
				"https://go.dev/doc/modules/developing",
			),
			Br(),
			P(`An <b>interface</b> describes behavior.`),

			quote(`Interfaces in Go provide a way to specify the
			  behavior of an object: if something can do this, then it
			  can be used here.`,
				"https://go.dev/doc/effective_go#interfaces_and_types",
			),
		),
	)
	p.NewCard(
		H3("module"),
	)
	p.NewCard(
		H3("package"),
	)
	p.NewCard(
		H3("interface"),
	)
	// ----------------------------------------

	p.NewCard(
		H2("Example"),
		deck.Middle(70,
			P(`Variations`),
		),
	)
	p.NewCard(
		H3("Direct"),
		deck.Middle(70,
			deck.Load("examples/direct.go"),
		),
	)
	p.NewCard(
		H3("Closure"),
		deck.Middle(70,
			deck.Load("examples/closure.go"),
		),
	)
	p.NewCard(
		H3("Method"),
		deck.Middle(70,
			deck.Load("examples/method.go"),
		),
	)
	p.NewCard(
		H3("Combo"),
		deck.Middle(70,
			deck.Load("examples/combo.go"),
			P(`More on the object oriented design side`),
		),
	)

	// ----------------------------------------

	p.NewCard(
		H2("Route"),
	)
	p.NewCard(
		H3("Naive"),
	)
	p.NewCard(
		H3("Combo"),
		TwoCol(
			deck.LoadLines("examples/combo_route.go", 9, -1),
			P(`Uses package http builtin singleton muxer.`),
			30, // right column width
		),
	)
	p.NewCard(
		H3("http/ServeMux"),
	)
	// ----------------------------------------
	p.NewCard(
		H2("Encode"),
	)
	p.NewCard(
		H3("encoding/json"),
	)
	// ----------------------------------------

	p.NewCard(
		H2("Test"),
		P(``),
	)
	p.NewCard(
		H3("http/httptest"),
	)
	p.NewCard(
		H3("coverage"),
	)

	// ----------------------------------------

	p.NewCard(
		H2("View"),
	)
	p.NewCard(
		H3("html/template"),
		P(``),
	)
	p.NewCard(
		H3("embed"),
	)

	p.Document().SaveAs("index.html")
}

func TwoCol(e1, e2 any, width int) *Element {
	leftWidth := 100 - 8 - width

	div := Div(Class("double"))
	div.With(
		Div(
			Class("column left"),
			Attr("style", fmt.Sprintf("width: %vvw", leftWidth)),
			e1,
		),
	)
	div.With(
		Div(
			Class("column right"),
			Attr("style", fmt.Sprintf("left: %vvw; width: %vvw", leftWidth+8, width)),
			e2,
		),
	)
	return div
}

func quote(el any, href string) *Element {
	return Quote(
		`"`, el, `"`,
		Br(),
		A(
			Href(href), strings.TrimPrefix(href, "https://"),
		),
	)
}
