package main

import (
	"fmt"
	"strings"

	. "github.com/gregoryv/web"
)

func main() {
	p := Deck{
		Title:     "Backend development",
		Author:    "Gregory Vinčić",
		AutoCover: true,
		cover:     Img(Src("img/coverpage.svg")),
		AutoTOC:   true,
	}
	// dark mode
	// p.Style("html, body, a", "background-color: #2E2E34", "color: #f0f8ff")
	p.Style(".header", "background-color: #f6f5f4",
		"border-bottom: 1px inset #000",
	)
	p.Style(".page .footer", "bottom: 1vh") // slightly up
	p.Style(".right>*, .right .shell",
		"margin-right: 1em",
	)

	// center toc if short titles
	p.Style(".toc",
		"font-size: 0.9em",
	)
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
		"display: block",
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
	p.Style(".stop>img",
		"text-align: center",
		"float:left",
		"margin-right: 20px",
	)
	p.Style(".stop>p",
		"padding-top: 50px",
		"color: red",
	)

	introSection(&p)
	conceptsSection(&p)
	httpExamples(&p)
	routeSection(&p)
	encodeSection(&p)
	testSection(&p)
	viewSection(&p)

	endSection(&p)
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
