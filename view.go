package main

import . "github.com/gregoryv/web"

func viewSection(p *Deck) {
	p.NewCard(
		H2("View"),

		Middle(60,
			P(`Serverside rendered HTML for several use cases`),
			Ul(
				Li("The obvious, you love the old school"),
				Li("API docs"),
				Li("Server stats"),
				Li("Performance, improve user experience and load times"),
			),
		),
	)
	p.NewCard(
		H3("html/template"),
		TwoCol(
			LoadEscaped("examples/htmltpl.go"),
			Wrap(
				LoadEscaped("examples/index.html"),
				ShellEscaped("$ go run examples/htmltpl.go", "examples/htmltpl.out"),
			),
			48,
		),
	)
	p.NewCard(
		H3("embed"),
	)
	// wip here
}
