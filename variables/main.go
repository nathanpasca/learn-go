package main

import "fmt"

var myAge int = 20 // can be inside or outside of function
var myFirstName = "Nathan" // if did not set the data types, the compiler will automatically set the datatypes based on the value.
const PI = 3.14 // declare a CONSTANT variable, constant variable means that the value inside the variable cannot be changed. And we must declare the value in the same time.
/* E.g const PI float32;
		PI = 3.15
		if we run it, it will result in an error, because we have to declare the value at the same line, and we must not change the value
*/

// we declare the name of the constant variable usually using UPPERCASE, to differentiate between a normal variable and a constant variable

var ( // We could make a multi function like this
	hisAge int = 22
	hisName string = "Resqi"
)

var a, b, c, d int = 1, 2, 3, 4 // or we could make a multiple function in one line
var e, f, g, h = 1, 2, "test", 3.14 // we could make the multiple function of different data types, we dont define the data type.

func main() {

	myLastName := "Pasca" // more simpler way to define or declare a variable but it cannot be declared outside of function

	fmt.Printf(myFirstName,myAge,myLastName)

}