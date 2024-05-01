package main

import (
	"fmt"
	"sync"
)

func main() {

	ch := make(chan int, 3)

	var wg sync.WaitGroup

	wg.Add(1)

	go func(ch chan int) {

		for i := 0; i < 10; i++ {

			ch <- 1

			fmt.Println("sending")

		}

		// recommended that only sender closes the channel not the receiver

		close(ch)

		wg.Done()

	}(ch)

	go func(ch chan int, wg *sync.WaitGroup) {

		for i := range ch { // receives values from the channel repeatedly until it is closed

			fmt.Printf("receiving: %v\n", i)

		}

		wg.Done()

	}(ch, &wg)

	wg.Wait()

	fmt.Println("main over-----------------------------")

}
