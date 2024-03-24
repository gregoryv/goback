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
	data, _ := json.Marshal(user)
	fmt.Print(string(data))
}

type Contact struct {
	Firstname string
	Lastname  string
}
