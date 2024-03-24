package main

import . "github.com/gregoryv/web"

func finallySection(p *Deck) {
	p.Style(".summary tr td:first-child",
		"padding-right: 2em",
	)
	p.Style(".summary tr th",
		"text-align: left",
		"border-bottom: 1px solid black",
	)
	p.NewCard(
		H2("Summary"),
		Middle(70,
			P(`Concepts covered in this presentation`),
			Table(Class("summary"),
				Tr(Th("Go language/structure", Attr("colspan", "2"))),
				row("module", "for dependency management"),
				row("package", "grouping related code"),
				row("interface", "behavior abstraction"),

				Tr(Th(Br(), "Package", Attr("colspan", "2"))),
				row("net/http", "handling requests and routing"),
				row("net/http/httptest", "utilities for HTTP testing"),
				row("testing", "automated testing of Go packages"),
				row("encoding/json", "marshaling datatypes to JSON"),
				row("html/template", "rendering HTML"),
				row("embed", "provides access to embedded files"),
			),
		),
	)
	p.Style(".courses",
		"width: 100%",
	)
	p.Style(".courses tr td",
		"vertical-align: top",
		"width: 33%",
	)
	p.Style(".course",
		"font-size: 1em",
		"width: 80%",
		"height: 90px",
		"border: 5px double black",
		"padding: 1em 1em",
		"text-align: center",
	)
	p.Style(".done",
		"color: #e2e2e2",
		"border: 5px double #e2e2e2",
	)
	p.Style(".maybe",
		"color: #e2e2e2",
	)
	p.NewCard(
		H3("Next time"),
		Br(),
		Table(Class("courses"),
			Tr(
				Td(
					Div(Class("course done"), "#1<br>Backend development"),
				),

				Td(
					Div(Class("course"), "#2<br>General problems"),
					Ol(
						Li("Test driven development"),
						Li("Package testdata"),
						Li("Documentation"),
						Li("Logging"),
						Li("Error handling"),
						Li("Database IO"),
					),
				),
				Td(
					Div(Class("course"), Em("#3<br>Concurrency problems")),
					Ol(Class("maybe"),
						Li("Concurrency"),
						Li("Benchmark"),
						Li("Testing strategy"),
						Li("... more"),
					),
				),
			),
		),
	)
}

func row(pkg, txt string) *Element {
	return Tr(
		Td(pkg),
		Td(txt),
	)
}
