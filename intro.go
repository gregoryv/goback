package main

import . "github.com/gregoryv/web"

func introSection(p *Deck) {
	p.NewCard(
		H2("Goal"),

		Middle(60,

			P(`The goal of this presentation is to teach you Go
		    concepts when writing programs serving HTTP
		    requests.`),

			P(`You will be learn about packages involved and their purpose.`),
		),
	)
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
			),
			40,
		),
	)
	p.NewCard(
		H2("Quick install"),
		Middle(70,
			P(`Download and install Go in the $HOME directory.`),
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
}
