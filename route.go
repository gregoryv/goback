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
		H3("naive"),
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
		H3("default http.ServeMux"),
		TwoCol(
			LoadLines("examples/singleton_mux.go", 9, -1),
			P(`Use http singleton ServeMux.`),
			30, // right column width
		),
	)

	p.NewCard(
		H3("your own http.ServeMux"),
		TwoCol(
			LoadFunc("examples/own_mux.go", "main"),
			P(`Create your own http.ServeMux.`),
			30, // right column width
		),
	)
}
