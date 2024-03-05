package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	if err := http.ListenAndServe(":8080", sayHello()); err != nil {
		log.Fatal(err)
	}
}

func sayHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, world!")
	}
}
