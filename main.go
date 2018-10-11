package main

import (
	"fmt"
	"moka/repl"
	"os"
)

func main() {
	fmt.Printf("Moka programming language v0.01\n")
	repl.Start(os.Stdin, os.Stdout)
}
