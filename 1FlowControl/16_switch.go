package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {

		case "java", "Java":
			fmt.Println("Extension is .java")

		case "c", "C":
			fmt.Println("Extension is .c")

		case "python", "Python":
			fmt.Println("Extension is .py")

		case "go", "Go":
			fmt.Println("Extension is .go")

		default:
			fmt.Println("Not valid option")

		}

		switch {
		case len(os.Args) == 2:
			fmt.Println("Only 1 arg. passed")

		case len(os.Args) > 2:
			fmt.Println("More than 1 arg. passed")
		}
	} else {
		fmt.Println("No arguments provided")
	}
}
