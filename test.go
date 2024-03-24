package main

import . "github.com/gregoryv/web"

func testSection(p *Deck) {
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
	// wip here
}
