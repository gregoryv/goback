package main

import (
	"fmt"
)

func main() {
	fmt.Println(painted(color(2)))
}

func painted(val fmt.Stringer) string {
	return fmt.Sprintf("%T.String(): %s", val, val.String())
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
