package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	if err := http.ListenAndServe(":8080", &Greeter{}); err != nil {
		log.Fatal(err)
	}
}

type Greeter struct{}

func (g *Greeter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}
