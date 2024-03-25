package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	people := []Person{
		// available people
		{Id: "p1", Name: "John"},
		{Id: "p2", Name: "Jane"},
	}
	mux := http.NewServeMux()
	mux.Handle("GET /person/{id}", servePersonById(people))

	// ... ListenAndServe omited for brevity
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
		// todo handle not found
	}
}

type Person struct {
	Id   string
	Name string
}
