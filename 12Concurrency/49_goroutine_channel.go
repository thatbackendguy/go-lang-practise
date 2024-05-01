package main

import (
	"fmt"
	"strings"
)

// declaring a function that computes the factorial of n

func factorial(n int, c chan int) {

	f := 1

	for i := 2; i <= n; i++ {

		f *= i

	}

	c <- f // sending the factorial value into the channel
}

func main() {

	ch := make(chan int) // declaring and initializing a channel of type `chan int`

	defer close(ch) // defer closing the channel

	go factorial(5, ch) // launching a goroutine

	// main() is waiting for a value to come from the channel

	// this is called a `blocking call`

	fmt.Println("5 factorial =", <-ch)

	// Spawning 20 goroutines that calculate the factorial

	for i := 1; i <= 20; i++ {

		go factorial(i, ch)

		fmt.Printf("Factorial of %d: %d\n", i, <-ch)

	}

	fmt.Println(strings.Repeat("#", 10))

	// Spawning another 10 goroutines this time as anonymous functions

	for i := 5; i < 15; i++ {

		go func(n int, c chan int) {

			f := 1

			for i := 2; i <= n; i++ {

				f *= i

			}

			// sending the value f into the channel

			c <- f
		}(i, ch)

		fmt.Printf("Factorial of %d is %d\n", i, <-ch)

	}

}
