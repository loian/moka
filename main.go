package main

import (
	"fmt"
	"os"
	"moka/repl"
)

func main() {
	fmt.Printf("Moka programming language v0.01\n")
	repl.Start(os.Stdin, os.Stdout)
}

