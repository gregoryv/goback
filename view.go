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

		TwoCol(
			LoadEscaped("examples/embed.go"),

			P(
				quote(
					"Package embed provides access to files embedded in the running Go program.",
					"https://pkg.go.dev/embed",
				),
			),
			40,
		),
	)
}
