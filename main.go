package main

import (
	"fmt"
	"os"
)

func main() {
	switch os.Args[1] {
	case "run":
		run()

	default:
		panic("bad command")
	}

	fmt.Println("Tente com 'axolotl' para rodar a CLI")
}

func run() {
	fmt.Println("Running the function :)")
}
