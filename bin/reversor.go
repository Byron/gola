package main

import (
	"fmt"
	"github.com/byron/gola"
	"github.com/xrash/smetrics"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "USAGE: %v <reverse-this>\n", os.Args[0])
		os.Exit(1)
	}

	fmt.Println(string.Reverse(os.Args[1]))
	s := smetrics.Jaro(os.Args[1], "Hello, World")
	switch {
	case s == 1.0:
		fmt.Println("BULLSEYE!")
	case s > 0.8:
		fmt.Println("Getting warmer")
	default:
		fmt.Println("Cold ...")
	}
}
