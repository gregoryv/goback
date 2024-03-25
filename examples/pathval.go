package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	// available people, ie. our data store
	people := []Person{
		{Id: "p1", Name: "John"},
		{Id: "p2", Name: "Jane"},
	}
	mux := http.NewServeMux()
	mux.Handle("GET /person/{id}", servePersonById(people))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

func servePersonById(people []Person) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		for _, person := range people {
			if person.Id == id {
				json.NewEncoder(w).Encode(person)
				return
			}
		}
		// ... handle not found
	}
}

type Person struct {
	Id   string
	Name string
}
