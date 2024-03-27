package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	handler := &Greeter{
		// complex setup
	}
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatal(err)
	}
}

type Greeter struct {
	// complex relations, eg. database
}

func (g *Greeter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}
