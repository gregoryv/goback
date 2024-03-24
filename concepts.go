package main

import . "github.com/gregoryv/web"

func conceptsSection(p *Deck) {
	p.NewCard(
		H2("Concepts"),
		Middle(70,

			P(`In Go; <b>modules</b> are used to group
			<b>packages</b> and functions.`),

			quote(`You can collect related packages into modules, then
		publish the modules for other developers to use.`,
				"https://go.dev/doc/modules/developing",
			),
			Br(),
			P(`An <b>interface</b> describes behavior.`),

			quote(`Interfaces in Go provide a way to specify the
			  behavior of an object: if something can do this, then it
			  can be used here.`,
				"https://go.dev/doc/effective_go#interfaces_and_types",
			),
		),
	)
	p.NewCard(
		H3("module"),
		TwoCol(
			Wrap(
				P(`Modules are initiated using the repository URL.`),
				Pre(Class("shell dark"),
					"$ go mod init github.com/gregoryv/uptime",
				),
				P(`This can then be used`),
				Pre(Class("shell dark"),
					"$ go get github.com/gregoryv/uptime[@VERSION]",
				),
				P("Find modules on ",
					A(Href("https://pkg.go.dev"), "pkg.go.dev"),
				),
			),
			Wrap(

				P(`Sharing code with others requires dependency
				management.`),

				pictogram("pic_module.png", "Pictogram of a Go module"),
			),
			50,
		),
	)
	p.NewCard(
		H3("package"),

		TwoCol(
			Wrap(
				// Span(Class("filename"), "myprogram.go"),
				Load("examples/imports.go"),
			),
			Wrap(
				Br(),
				quote(
					Wrap(
						`... package name should be good: short, concise, evocative. `,
						"...",
						`Use the package structure to help you choose good names. `,
					),

					"https://go.dev/doc/effective_go#package-names",
				),
				pictogram("pic_package.png", "Pictogram of a Go package"),
			),
			50,
		),
	)
	p.NewCard(
		H3("interface"),

		TwoCol(
			Wrap(
				Br(),
				quote(
					`Interfaces with only one or two methods are common in Go
			code, and are usually given a name derived from the
			method, such as io.Writer for something that implements
			Write.`,

					"https://go.dev/doc/effective_go#interfaces_and_types",
				).With(Class("small")),

				Load("examples/stringer.go"),
				pictogram("pic_interface.png", "Pictogram of a Go interface"),
			),
			LoadLines("examples/interface.go", 7, -1),
			50,
		),
	)
}
