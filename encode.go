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
	p.NewCard(
		H3("nice json"),
		TwoCol(
			Load("examples/nicejson.go"),
			Shell("$ go run ./examples/nicejson.go", "./examples/nicejson.json"),
			40,
		),
	)
	p.NewCard(
		H3("field tags"),
		TwoCol(
			Load("examples/fieldtags.go"),
			Wrap(
				P("Control naming of fields, eg. lowercase names."),
				Shell("$ go run ./examples/fieldtags.go", "./examples/fieldtags.json"),
			),
			40,
		),
	)
}
