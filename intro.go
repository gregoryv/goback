package main

import . "github.com/gregoryv/web"

func introSection(p *Deck) {
	p.NewCard(
		H2("Introduction"),

		Middle(60,

			P(`The goal of this presentation is to teach you Go
		    concepts when writing applications serving HTTP
		    requests.`),
		),
	)
	p.NewCard(
		H3("History"),
		Img(Src("img/people.png")),
	)
	p.NewCard(
		H3("Quick install"),
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
