package main

import . "github.com/gregoryv/web"

func finallySection(p *Deck) {
	p.Style(".summary tr td:first-child",
		"padding-right: 2em",
	)
	p.NewCard(
		H2("Summary"),
		TwoCol(
			Wrap(
				P(`Key knowledge areas shared in this presentation`),
				Table(Class("summary"),
					Tr(Th("Package")),
					row("net/http", "handling requests and routing"),
					row("net/http/httptest", "utilities for HTTP testing"),
					row("testing", "automated testing of Go packages"),
					row("encoding/json", "marshaling datatypes to JSON"),
					row("embed", "provides access to embedded files"),
				),
			),
			Wrap(
				H4("Concepts"),
				P("Design variations related to request handling"),
			),
			30,
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
