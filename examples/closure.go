package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":8080", sayHello())
	if err != nil {
		log.Fatal(err)
	}
}

func sayHello() http.HandlerFunc {
	txt := "Hello, world!"
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, txt)
	}
}
