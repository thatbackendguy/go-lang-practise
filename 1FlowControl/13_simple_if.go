package main

import (
	"fmt"
	"os"
)

func main() {

	// i and err are variables scoped to the if statement only
	if args := os.Args; len(os.Args) > 1 {
		fmt.Println("1st arg: ", args[1])
	} else {
		fmt.Println("no args passed")
	}
}
