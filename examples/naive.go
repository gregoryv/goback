package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	var ctl Controller
	if err := http.ListenAndServe(":8080", &ctl); err != nil {
		log.Fatal(err)
	}
}

func (c *Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if r.URL.Path == "/" {
			c.sayHello(w, r)
			return
		}
		if r.URL.Path == "/bye" {
			c.sayGoodbye(w, r)
			return
		}
	}
	http.Error(w, "routing failed", http.StatusBadRequest)
}

type Controller struct{}

func (c *Controller) sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

func (c *Controller) sayGoodbye(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Goodbye, world!")
}
