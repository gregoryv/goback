package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	h := &Hotel{
		Name: "Purple glow",
		Rooms: []Room{
			{Number: 1},
			{Number: 2},
			{Number: 3},
		},
	}

	mux := http.NewServeMux()
	mux.Handle("POST /room/{number}", BookRoom(h.Rooms))
	mux.Handle("GET /room", ServeRooms(h.Rooms))
	mux.Handle("GET /", ServeHotel(h))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

type Hotel struct {
	Name  string `json:"name"`
	Rooms []Room `json:"-"` // rooms are rendered separately
}

type Room struct {
	Number string `json:"number"` // using string for simplicty
	Booked bool   `json:"booked"`
}

func BookRoom(rooms []Room) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		number := r.PathValue("number")
		// update room state
		for i := range rooms {
			if i != number {
				continue
			}
			rooms[i].Booked = true
		}
		json.NewEncoder(w).Encode(rooms)
	}
}

func ServeRooms(rooms []Room) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(rooms)
	}
}

func ServeHotel(h *Hotel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(h)
	}
}
