package main

import . "github.com/gregoryv/web"

func httpExamples(p *Deck) {
	p.NewCard(
		H2("HTTP"),
		Middle(68,

			P(`Respond to HTTP requests in Go by implementing the
			http.Handler interface.`),

			Load("examples/handler.go"),
		),
	)
	p.NewCard(
		H3("Direct"),
		TwoCol(
			Load("examples/direct.go"),
			Wrap(

				P(`Direct design uses named functions as http
                handlers, implementing the http.HandlerFunc
                interface`),

				pictogram("pic_func.png", "Func"),

				Span(Class("small"),
					Pre(Class("shell dark"),
						"$ go run helloworld.go",
					),
				),
			),
			40,
		),
	)
	p.NewCard(
		H3("Closure"),
		TwoCol(
			Load("examples/closure.go"),
			Wrap(

				P(`Closure design uses named functions returning a
                http.Handler.`),

				pictogram("pic_closure.png", "Closure"),
			),
			40,
		),
	)
	p.NewCard(
		H3("Method"),
		TwoCol(
			Load("examples/method.go"),
			Wrap(

				P(`Method design declares types that implement
				http.Handler`),

				pictogram("pic_method.png", "Method"),
			),
			40,
		),
	)
	p.NewCard(
		H3("Combo"),
		TwoCol(
			Span(Class("small"),
				Load("examples/combo.go"),
			),
			Wrap(

				P(`Combo design declares types with multiple methods
				that are either methods that either have a
				http.HandlerFunc signature or return one (closure).`),

				pictogram("pic_combined_closure.png", "Combo"),
			),
			45,
		),
	)
}
