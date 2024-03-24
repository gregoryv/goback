package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	user := Contact{
		Firstname: "John",
		Lastname:  "Doe",
		Address: &Address{
			Street: "Lexington road",
			Number: 137,
		},
	}
	data, _ := json.MarshalIndent(user, "", "  ")
	fmt.Print(string(data))
}

type Contact struct {
	Firstname string `json:"firstname"` // change name
	Lastname  string `json:"-"`         // ignore this field
	*Address  `json:"address"`
}

type Address struct {
	Street     string `json:"street"`
	Number     uint   `json:"no"`                    // rename entirely
	PostalCode int    `json:"postal_code,omitempty"` // skip empty value
}
