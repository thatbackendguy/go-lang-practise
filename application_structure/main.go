package main

import "fmt"

const secondsInHour = 3600

func main() {
	fmt.Println("Hello World")

	distance := 60.8 // distance in KMs

	fmt.Printf("The distance in miles: %f\n", distance*0.621)
	fmt.Println(secondsInHour)
}
