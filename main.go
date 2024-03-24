package main

import (
	"fmt"
	"strings"

	. "github.com/gregoryv/web"
)

func main() {
	p := Deck{
		Title:     "Golang; Backend development",
		Author:    "Gregory Vinčić",
		AutoCover: true,
		AutoTOC:   true,
	}
	// dark mode
	// p.Style("html, body, a", "background-color: #2E2E34", "color: #f0f8ff")
	p.Style(".header", "background-color: #f6f5f4",
		"border-bottom: 1px inset #000",
	)
	p.Style(".page .footer", "bottom: 1vh") // slightly up
	// center toc if short titles
	// p.Style(".toc", "margin-left: 10vw", "width: 20vw")
	p.Style(".shell",
		"padding: 1em",
		"border-radius: 10px",
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
		"font-size: 1.5vw",
	)
	p.Style("p>a",
		"text-decoration: underline",
	)
	p.Style("quote",
		"display: inline-block",
		"font-style: italic",
		"padding-left: 2vw",
		"padding-right: 2vw",
		"font-size: 1.5vw",
	)
	p.Style("quote>a",
		"font-size: 1vw",
		"float: right",
	)
	p.Style("quote.small",
		"font-size: 1vw",
	)

	p.Style(".pictogram",
		"width: 400px",
		"text-align: center",
		"padding: 1em 1em",
		"margin: 1em auto",
	)
	p.Style(".pictogram label",
		"font-style: italic",
		"font-size: 0.6em",
	)
	p.Style(".srcfile",
		"font-size: 1.3vw",
	)
	p.Style(".filename",
		"display: block",
		"font-size: 0.8vw",
		"text-align: right",
		"margin-top: 2em",
		"margin-bottom: -2em",
	)
	p.Style(".icons",
		"column-count: 3",
		"list-style-type: none",
	)

	p.Style(".small *",
		"font-size: 0.9vw",
	)
	p.Style(".small .filename",
		"margin-bottom: -1.5em",
	)
	p.Style(".right>*",
		"margin-right: 1em",
	)
	p.Style(".stop>img",
		"text-align: center",
		"float:left",
		"margin-right: 20px",
	)
	p.Style(".stop>p",
		"padding-top: 50px",
		"color: red",
	)
	// wip style
	// ----------------------------------------  ----------------------------------------
	p.NewCard(
		H2("Quick start"),
		Middle(70,
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

	conceptsSection(&p)
	httpExamples(&p)
	routeSection(&p)
	encodeSection(&p)
	testSection(&p)
	viewSection(&p)

	// ----------------------------------------

	p.NewCard(
		H2("Next time"),

		Pre(`
Design patterns
  - access control

Benchmarking
Logging
Error handling

`),
	)
	p.Document().SaveAs("index.html")
}

func stop(txt string) *Element {
	return Div(Class("stop"),
		Img(Src("img/stop.png")),
		P(txt),
	)
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
		Br(),
		A(
			Href(href), strings.TrimPrefix(href, "https://"),
		),
	)
}

func pictogram(filename, label string) *Element {
	return Div(
		Class("pictogram"),
		Img(Src("img/"+filename)),
		Br(),
		Label(label),
	)
}
