package main

import (
	"bibi/repl"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Bibi's Incredibly Basic Interpreter ;)\n")
	fmt.Printf("Write something awesome!\n")
	repl.Start(os.Stdin, os.Stdout)
}
