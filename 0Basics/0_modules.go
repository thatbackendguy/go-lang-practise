package main

import (
	"fmt"
	"github.com/thatbackendguy/go-lang-practise/8Functions"
)

func main() {
	fmt.Println(_Functions.MyName + "@" + _Functions.GetIp())
	fmt.Println(_Functions.GetQuote())
}
