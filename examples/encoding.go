package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	user := Contact{
		Firstname: "John",
		Lastname:  "Doe",
		Age:       23,
		Alive:     true,
	}
	data, _ := json.Marshal(user)
	fmt.Print(string(data))
}

type Contact struct {
	Firstname string
	Lastname  string
	Age       int
	Alive     bool
}
