package main

import (
	"fmt"
	"math"
)

func main() {
	var x uint8 = 255

	x++ // overflow, no panic thrown, as operation is done at runtime

	fmt.Println(x)

	//a := int8(255 + 1)
	//fmt.Println(a)

	var b int8 = 127

	fmt.Printf("%d\n", b+1)

	b = -128

	b--

	fmt.Printf("%d\n", b)

	f := float32(math.MaxFloat32)

	fmt.Printf("%f\n", f)

	f = f * 1.2

	fmt.Printf("%f\n", f)

	name := "yash"

	var heart int = 10084

	ans := string(heart) + string(32) + name + " hello"

	fmt.Printf("%s\n", ans)

}
