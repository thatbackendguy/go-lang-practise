package main

import "fmt"

//** COMMENTS **//

// this is a single line comment

/*
 This is a block comment.
 a := 10
 fmt.Println(a)
*/

var name = "John Wick" // inline comment

//** NAMING CONVENTIONS IN GO **//

// Naming Conventions are important for code readability and maintainability.

// use short, concise names especially in shorter scopes
// common names for common types:
var s string   //string
var i int      //index
var num int    //number
var msg string //message
var v string   //value
var err error  //error value
var done bool  //bool, has been done?

// use mixedCase a.k.a camelCase instead of snake_case (variables and  functions)
var maxValue = 100  // recommended (camelCase)
var max_value = 100 // not recommended (snake_case)

// recommended
func writeToFile() {
}

// not recommended
func write_to_file() {
}

// write acronyms in all caps
var writeToDB = true // recommended
var writeToDb = true // not recommended

func main() {
	// using var keyword
	// keyword identifier type
	var age int = 22

	fmt.Println("Age: ", age)

	// compile time error if there is un-used variable
	// It's not mandatory to use a variable declared at package scope. Only local variables must be used.

	var name = "Yash" // type inference

	_ = name // Blank identifier: used to remove the error for unused variables

	// using shorthand: works only in block scope
	// cannot use shorthand declaration for package scoped variables
	s := "Learning Go-lang!"

	fmt.Println(s)

	// Vid-16. MULTIPLE DECLARATIONS
	car, cost := "Audi", 100000

	car, year := "BMW", 2024

	fmt.Println(car, cost, year)

	var opened = false

	opened, file := true, "a.txt"

	_, _ = opened, file

	var (
		salary    float64 // 0
		firstName string  // ""
		gender    bool    // false
	)

	fmt.Println(salary, firstName, gender)

	var a, b, c int

	a, b, c = 1, 2, 3

	fmt.Println(a, b, c)

	var i, j int

	i, j = 10, 20

	i, j = j, i // swapping variables

	fmt.Println(i, j)

	var surname = "Prajapati"

	fullName := surname + " " + name

	fmt.Println(fullName)

	/*
		KEYWORDS (25): cannot be used as identifier
		break     default      func    interface  select
		case      defer        go      map        struct
		chan      else         goto    package    switch
		const     fallthrough  if      range      type
		continue  for          import  return     var
	*/
	const a1 = 5.0
	const name1 = 2.3 + 5
	_ = name

	// use fewer letters, don’t be too verbose especially in smaller scopes
	var packetsReceived int // NOT OK, too verbose

	var n int // OK

	_, _ = packetsReceived, n

	// an uppercase first letter has special significance to go (it will be exported in other packages)

}
