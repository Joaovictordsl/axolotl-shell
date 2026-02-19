package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("test")
	for {
		fmt.Print(">")

		scanner.Scan()

		line := scanner.Text()
		if line == "exit" {
			break
		}
		fmt.Println(line)
	}
	if err := scanner.Err(); err != nil {

		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
}
