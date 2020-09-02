package main

import (
	"fmt"
	"os"

	"expensy/cmd/expensyctl"
)

func main() {
	cmd := os.Args[1]
	args := os.Args[2:]
	length := len(args)

	switch cmd {
	case "add":
		fmt.Println(cmd)
		addCommand(args)
	default:
		fmt.Println("Default case")
	}
}
