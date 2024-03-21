package main

import (
	"fmt"
	"time"

	"github.com/gregoryv/uptime"
)

func main() {
	start := time.Now()
	dur := uptime.Since(start)
	fmt.Println(dur.String())
}
