package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	var ctl Controller

	if err := http.ListenAndServe(":8080", ctl.sayHello()); err != nil {
		log.Fatal(err)
	}
}

type Controller struct{}

func (c *Controller) sayHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, world!")
	}
}
