package main

import (
	_ "embed" // used at compile time
	"fmt"
)

func main() {
	fmt.Print(index)
}

//go:embed index.html
var index string
