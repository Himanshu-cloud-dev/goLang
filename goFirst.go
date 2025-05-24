package main // The entry of any go programm.

import "fmt" // importing fmt package.

var firstName  = "Himanshu"

type Employee struct {
	FirstName string
	LastName string
}

func main() {
	// naming convension := // pascal case ( Structs ,  Interfaces , enums)
	// snake_case 
	// UPPERCASE
	// mixedCase

	const userLife int = 5

	var name string = "Himanshu"
	lastName := "Bhardwaj"
	firstName = name
	fmt.Println("Hello, " + name + " " + lastName)  // function provided by the fmt package.
	printName()
	} //  Entry point of any go Programm.

// go build <fileName> :- with this command we can convert our go programe in standalone binary format.

func printName () {
	fmt.Println(firstName )
}