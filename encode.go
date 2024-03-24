package main

import (
	. "github.com/gregoryv/web"
)

func encodeSection(p *Deck) {
	p.NewCard(
		H2("Encode"),
		Middle(60,
			P("From the dictionary"),

			Pre(`encode
  
    &lt;algorithm, hardware&gt; To convert {data} or some physical
    quantity into a given format.  E.g. {uuencode}.
`),

			P("In Go this means convert ", Code("type X"), " to ", Code("[]byte")),
		),
	)

	p.NewCard(
		H3("encoding/json"),
		TwoCol(
			Load("examples/encoding.go"),
			Shell("$ go run ./examples/encoding.go", "./examples/encoding.json"),
			40,
		),
	)

	// wip here
}
