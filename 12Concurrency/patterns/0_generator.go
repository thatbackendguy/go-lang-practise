package main

import "fmt"

func Generator(message string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 100; i++ {
			c <- fmt.Sprintf("%s : %d", message, i)
		}
		close(c)
	}()
	return c
}

//func main() {
//	yash := Generator("yash")
//	shivam := Generator("shivam")
//
//	for i := 0; i < 10; i++ {
//		fmt.Println(<-yash)
//		fmt.Println(<-shivam)
//	}
//
//}
