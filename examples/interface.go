package main

import (
	"fmt"
)

func main() {
	fmt.Println(painted(color(2)))
}

func painted(value fmt.Stringer) string {
	return fmt.Sprintf("Painted: %s", value.String())
}

type color int

func (c color) String() string {
	switch c {
	case 1:
		return "red"
	case 2:
		return "green"
	default:
		return "white"
	}
}
