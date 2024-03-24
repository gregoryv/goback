package main

import . "github.com/gregoryv/web"

func finallySection(p *Deck) {
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
	// wip here
}
