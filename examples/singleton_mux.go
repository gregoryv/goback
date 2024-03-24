package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	var ctl Controller
	http.HandleFunc("GET /", ctl.sayHello)
	http.HandleFunc("GET /bye", ctl.sayGoodbye)
	if err := http.ListenAndServe(":8080", nil); err != nil {
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
