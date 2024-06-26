package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	user := Contact{
		Firstname: "John",
		Lastname:  "Doe",
	}
	data, _ := json.MarshalIndent(user, "", "  ")
	fmt.Print(string(data))
}

type Contact struct {
	Firstname string `json:"firstname"` // change name
	Lastname  string `json:"-"`         // ignore this field
}
