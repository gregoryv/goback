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
		TwoCol(
			Load("examples/httptest/main_test.go"),
			Wrap(
				P(`Quick unit test of specific handlers using httptest.Recorder`),
				Ol(
					Li("setup"),
					Li("call handler"),
					Li("check response"),
				),
			),
			40,
		),
	)
	p.NewCard(
		H3("Coverage"),
		Shell("$ go test -v -cover .", "examples/httptest_cover.out"),
		Div(Class("small"),
			TwoCol(
				Wrap(
					Div(Class("filename"), "main_test.go"),
					LoadLines("examples/httptest/main_test.go", 9, -1),
				),
				Wrap(
					Div(Class("filename"), "main.go"),
					LoadLines("examples/httptest/main.go", 9, -1),
				),
				50,
			),
		),
	)
}
