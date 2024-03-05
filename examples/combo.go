package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	var ctl Controller

	if err := http.ListenAndServe(":8080", http.HandlerFunc(ctl.sayHello)); err != nil {
		log.Fatal(err)
	}
}

type Controller struct{}

func (c *Controller) sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}
