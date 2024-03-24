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

	p.NewCard(
		H3("Next time"),

		Pre(`
Design patterns
  - access control

Benchmarking
Logging
Error handling

`),
	)
	// wip here
}

func row(pkg, txt string) *Element {
	return Tr(
		Td(pkg),
		Td(txt),
	)
}
