package main

import . "github.com/gregoryv/web"

func introSection(p *Deck) {
	/*p.NewCard(
		H2("Goal"),

		Middle(60,

			P(`The goal of this presentation is to teach you Go
		    concepts when writing programs serving HTTP
		    requests.`),

			P(`You will be learn about packages involved and their purpose.`),
		),
	)*/
	p.Style(".narrow",
		"display: block",
		"margin: auto auto",
		"width: 70%",
	)
	p.NewCard(
		H2("History"),

		Div(Class("narrow"),
			P(`Robert Griesemer, Ken Thompson and Rob Pike are the
		original authors of the Go language which is currently
		supervised by Russ Cox. Ian wrote the first compiler as a
		frontend to gcc.`),
		),
		TwoCol(
			Center(Div(Class("small"),
				Pre(`The story

  1960       1970          1980          1990        2000      2007-2009

                                                +- Javascript  --+
                                               /                  \
                +----- C -----------------+---+ PHP/Perl        ---+
               /                           \                        \
- Algol ------+                             +-- Java             ----+- Go
               \                                                    /
                \       +--- Modula ---+-- Python               -- +
                 \     /                                          /
                  +---+----------+         o-- Ada 95            /
                  |               \                             /
              Pascal               +---- Oberon             ---+
`),
			),
			),

			Wrap(
				Br(),

				Img(Src("img/people.png")),

				P(`They needed a language for writing servers on
				modern multicore computers`),
			),
			40,
		),
	)
	p.NewCard(
		H2("Quick install"),
		Middle(70,
			P(`Download and install Go in the $HOME directory.`),
			Center(A(Href("https://go.dev/dl"), "https://go.dev/dl")),

			Pre(Class("shell dark"),
				`$ cd $HOME
$ wget https://go.dev/dl/go1.22.0.linux-amd64.tar.gz
$ tar xvfz go1.22.0.linux-amd64.tar.gz
$ export PATH="$PATH:$HOME/go/bin"
$ go version
go version go1.22.0 linux/amd64`,
			),

			P("For more information refer to ",
				A(Href("https://go.dev/doc/install"),
					`"Download and install"`),
				".",
			),
		),
	)
	p.Style(".structures",
		"width: 100%",
	)
	p.Style(".structures td",
		"vertical-align: top",
		"padding-right: 1em",
	)
	p.Style(".structures td *",
		"font-size: 0.7em",
	)
	p.Style(".structures, td.wide",
		"width: 100%",
	)

	p.NewCard(
		H2("Project structure"),

		Middle(80,

			Table(Class("structures wide"),
				Tr(
					Td(
						Shell("$ cd neon; tree/", "examples/neon.tree"),
						quote(
							"A basic Go package has all its code in the project’s root directory.",
							"https://go.dev/doc/modules/layout",
						),
					),
					Td(
						Shell("$ cd neon/; tree", "examples/neon2.tree"),
					),
					Td(
						Shell("$ cd neon/; tree", "examples/neon3.tree"),

						quote("Server project",
							"https://go.dev/doc/modules/layout#server-project",
						),
					),
					Td(Class("wide"),
						Shell("$ cd neon/; tree", "examples/neon4.tree"),

						quote("Don't waste a good name",
							"Gregory Vinčić",
						),
					),
				),
			),
		),
	)
}
