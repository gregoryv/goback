package main

import . "github.com/gregoryv/web"

func viewSection(p *Deck) {
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
	// wip here
}
