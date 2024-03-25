package main

import . "github.com/gregoryv/web"

func routeSection(p *Deck) {
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

	p.NewCard(
		H3("Default http.ServeMux"),
		TwoCol(
			LoadLines("examples/singleton_mux.go", 9, -1),
			P(`Use http singleton ServeMux.`),
			35, // right column width
		),
	)

	p.NewCard(
		H3("Your own http.ServeMux"),
		TwoCol(
			LoadLines("examples/own_mux.go", 9, -1),
			P(`Create your own http.ServeMux.`),
			37, // right column width
		),
	)
	p.NewCard(
		H3("Path value"),
		TwoCol(
			LoadLines("examples/pathval.go", 1, 21),
			LoadLines("examples/pathval.go", 23, -1),
			46,
		),
	)
}
