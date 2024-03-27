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

			P("In Go this equals to marshal ", Code("type X"), " to ", Code("[]byte")),
		),
	)

	p.NewCard(
		H3("encoding/json"),
		TwoCol(
			Load("examples/encoding.go"),
			Shell("$ go run ./examples/encoding.go", "./examples/encoding.json"),
			50,
		),
	)
	p.NewCard(
		H3("Nice json"),
		TwoCol(
			Load("examples/nicejson.go"),
			Shell("$ go run ./examples/nicejson.go", "./examples/nicejson.json"),
			40,
		),
	)
	p.NewCard(
		H3("Struct tags"),
		TwoCol(
			Load("examples/fieldtags.go"),
			Wrap(
				P("Control naming of fields, eg. lowercase names."),
				Shell("$ go run ./examples/fieldtags.go", "./examples/fieldtags.json"),
			),
			40,
		),
	)
	p.NewCard(
		H3("Complex struct"),
		TwoCol(
			LoadLines("examples/complex_struct.go", 20, -1),
			Wrap(
				Shell("$ go run ./examples/complex_struct.go", "./examples/complex_struct.json"),
			),
			38,
		),
	)
}
