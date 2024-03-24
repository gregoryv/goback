package main

import (
	. "github.com/gregoryv/web"
	"github.com/gregoryv/web/files"
)

func testSection(p *Deck) {
	p.NewCard(
		H2("Test"),
		TwoCol(
			Wrap(
				Div(Class("filename"), "main_test.go"),
				Load("examples/project/main_test.go"),

				Div(Class("filename"), "main.go"),
				LoadLines("examples/project/main.go", 6, -1),
			),
			Wrap(
				Pre(
					files.MustLoad("examples/project.tree"),
				),
				Shell("$ go test -v .", "examples/test_project.out"),
			),
			50,
		),
	)
	p.NewCard(
		H3("http/httptest"),
	)
	p.NewCard(
		H3("coverage"),
	)
	// wip here
}
