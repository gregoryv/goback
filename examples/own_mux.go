package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	var ctl Controller

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", ctl.sayHello)
	mux.HandleFunc("GET /bye", ctl.sayGoodbye)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

type Controller struct{}

func (c *Controller) sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

func (c *Controller) sayGoodbye(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Goodbye, world!")
}
