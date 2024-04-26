package main

import (
	"fmt"
	"github.com/thatbackendguy/go-lang-practise/Functions"
)

func main() {
	fmt.Println(Functions.MyName + "@" + Functions.GetIp())
	fmt.Println(Functions.GetQuote())
}
