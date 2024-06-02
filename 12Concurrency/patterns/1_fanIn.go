package main

//
//import (
//	"fmt"
//	"runtime"
//)
//
//func fanIn(inpu1, inpu2 <-chan string) <-chan string {
//	c := make(chan string)
//
//	go func() {
//		for {
//			c <- <-inpu1
//		}
//	}()
//	go func() {
//		for {
//			c <- <-inpu2
//		}
//	}()
//	return c
//}
//
//func main() {
//	c := fanIn(Generator("yash"), Generator("shivam"))
//
//	fmt.Println("Current running go-routines:", runtime.NumGoroutine())
//
//	for i := 0; i < 30; i++ {
//		fmt.Println(<-c)
//	}
//
//	fmt.Println("Exiting main!")
//}
