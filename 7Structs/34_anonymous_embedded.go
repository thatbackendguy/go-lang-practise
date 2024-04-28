package main

import (
	"fmt"
	"strings"
)

func main() {

	// an anonymous struct is a struct with no explicitly defined struct type alias.

	yash := struct {
		firstName, lastName string

		age int
	}{

		firstName: "Yash",

		lastName: "Prajapati",

		age: 21,
	}

	fmt.Printf("%#v\n", yash)

	// =>struct { firstName string; lastName string; age int }{firstName:"Yash", lastName:"Prajapati", age:21

	//** ANONYMOUS FIELDS **//

	// fields type becomes fields name.

	type Book struct {
		string

		float64

		bool
	}

	b1 := Book{"Meluha by Amish Tripathi", 204.99, false}

	fmt.Printf("%#v\n", b1) // => main.Book{string:"Meluha by Amish Tripathi", float64:204.99, bool:false}

	fmt.Println(b1.string) // => Meluha by Amish Tripathi

	// mixing anonymous with named fields:

	type Employee1 struct {
		name string

		salary int

		bool
	}

	e := Employee1{"Yash", 40000, false}

	fmt.Printf("%#v\n", e) // => main.Employee1{name:"Yash", salary:40000, bool:false}

	e.bool = true // changing a field

	fmt.Println(strings.Repeat("#", 10))

	//** EMBEDDED STRUCTS **//

	// An embedded struct is just a struct that acts like a field inside another struct.

	// define a new struct type

	type Contact struct {
		email, address string

		phone int
	}

	// define a struct type that contains another struct as a field

	type Employee struct {
		name string

		salary int

		contactInfo Contact
	}

	// declaring a value of type Employee

	yash1 := Employee{

		name: "Yash P",

		salary: 40000,

		contactInfo: Contact{

			email: "yashpra1010@gmail.com",

			address: "Gujarat, India",

			phone: 987654,
		},
	}

	fmt.Printf("%+v\n", yash1)

	// => {name:Yash P salary:40000 contactInfo:{email:yashpra1010@gmail.com address:Gujarat, India phone:987654}}

	// accessing a field

	fmt.Printf("Employee's salary: %d\n", yash1.salary) // => Employee's salary: 40000

	// accessing a field from the embedded struct

	fmt.Printf("Employee's email:%s\n", yash1.contactInfo.email) // => Employee's email:yashpra1010@gmail.com

	// updating a field

	yash1.contactInfo.email = "yash@thatbackendguy.com"

	fmt.Printf("Employee's new email address:%s\n", yash1.contactInfo.email)

	// => Employee's new email address:yash@thatbackendguy.com

}
