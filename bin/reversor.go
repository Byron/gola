package main

import (
	"fmt"
	"github.com/byron/gola"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "USAGE: %v <reverse-this>\n", os.Args[0])
		os.Exit(1)
	}

	fmt.Println(string.Reverse(os.Args[1]))
}
