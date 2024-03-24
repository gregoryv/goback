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
	p.Style(".right .srcfile, .right .shell, .right p",
		"margin-right: 1em",
	)
	p.Style(".stop img",
		"text-align: middle",
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

	// ----------------------------------------
	p.NewCard(
		H2("Concepts"),
		Middle(70,

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
		TwoCol(
			Wrap(
				P(`Modules are initiated using the reposity URL.`),
				Pre(Class("shell dark"),
					"$ go mod init github.com/gregoryv/uptime",
				),
				P(`This can then be used`),
				Pre(Class("shell dark"),
					"$ go get github.com/gregoryv/uptime[@VERSION]",
				),
				P("Find modules on ",
					A(Href("https://pkg.go.dev"), "pkg.go.dev"),
				),
			),
			Wrap(

				P(`Sharing code with others requires dependency
				management.`),

				pictogram("pic_module.png", "Pictogram of a Go module"),
			),
			50,
		),
	)
	p.NewCard(
		H3("package"),

		TwoCol(
			Wrap(
				// Span(Class("filename"), "myprogram.go"),
				Load("examples/imports.go"),
			),
			Wrap(
				Br(),
				quote(
					Wrap(
						`... package name should be good: short, concise, evocative. `,
						"...",
						`Use the package structure to help you choose good names. `,
					),

					"https://go.dev/doc/effective_go#package-names",
				),
				pictogram("pic_package.png", "Pictogram of a Go package"),
			),
			50,
		),
	)
	p.NewCard(
		H3("interface"),

		TwoCol(
			Wrap(
				Br(),
				quote(
					`Interfaces with only one or two methods are common in Go
			code, and are usually given a name derived from the
			method, such as io.Writer for something that implements
			Write.`,

					"https://go.dev/doc/effective_go#interfaces_and_types",
				).With(Class("small")),

				Load("examples/stringer.go"),
				pictogram("pic_interface.png", "Pictogram of a Go interface"),
			),
			LoadLines("examples/interface.go", 7, -1),
			50,
		),
	)
	// ----------------------------------------

	p.NewCard(
		H2("HTTP Examples"),
		Middle(68,

			P(`Respond to HTTP requests in Go by implementing the
			http.Handler interface.`),

			Load("examples/handler.go"),
		),
	)
	p.NewCard(
		H3("Direct"),
		TwoCol(
			Load("examples/direct.go"),
			Wrap(

				P(`Direct design uses named functions as http
                handlers, implementing the http.HandlerFunc
                interface`),

				pictogram("pic_func.png", "Func"),

				Span(Class("small"),
					Pre(Class("shell dark"),
						"$ go run helloworld.go",
					),
				),
			),
			40,
		),
	)
	p.NewCard(
		H3("Closure"),
		TwoCol(
			Load("examples/closure.go"),
			Wrap(

				P(`Closure design uses named functions returning a
                http.Handler implementation`),

				pictogram("pic_closure.png", "Closure"),
			),
			40,
		),
	)
	p.NewCard(
		H3("Method"),
		TwoCol(
			Load("examples/method.go"),
			Wrap(

				P(`Method design declares types that implement
				http.Handler`),

				pictogram("pic_method.png", "Method"),
			),
			40,
		),
	)
	p.NewCard(
		H3("Combo"),
		TwoCol(
			Span(Class("small"),
				Load("examples/combo.go"),
			),
			Wrap(

				P(`Combo design declares types with multiple methods
				that are either methods that either have a
				http.HandlerFunc signature or return one (closure).`),

				pictogram("pic_combined_closure.png", "Combo"),
			),
			45,
		),
	)

	// ----------------------------------------
	p.NewCard(
		H2("Route"),

		Middle(68,

			P(`For systems serving many resources`),

			Load("examples/servemux.go"),

			Pre(
				`multiplexer
      n 1: a device that can interleave two or more activities
`,
			),
		),
	)
	p.NewCard(
		H3("Naive"),
		TwoCol(
			LoadLines("examples/naive.go", 16, 27),
			Wrap(
				P(`Manual routing.`),
				stop(`Error prone and complex.`),
			),
			37, // right column width
		),
	)
	// wip here
	p.NewCard(
		H3("Combo"),
		TwoCol(
			LoadLines("examples/combo_route.go", 9, -1),
			P(`Use http singleton ServeMux.`),
			30, // right column width
		),
	)

	p.NewCard(
		H3("http.ServeMux"),
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
