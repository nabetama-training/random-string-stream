package main

import (
	"flag"
	"fmt"

	"github.com/nabetama-training/random-string-stream"
)

var (
	n = flag.Int("n", 0, "N strings")
)

func main() {
	flag.Parse()
	if *n == 0 {
		for s := range randomString.Stream() {
			fmt.Print(s)
		}
	} else {
		for i := 0; i < *n; i++ {
			fmt.Print(randomString.String())
		}
	}
	fmt.Println()
}
