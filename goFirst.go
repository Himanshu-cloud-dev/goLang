package main // The entry of any go programm.

import "fmt" // importing fmt package.

var firstName  = "Himanshu"

func main() {
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