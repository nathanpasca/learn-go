package main

import "fmt"

var a string = "Nathan"
var b, c int = 20, 21
var d string = "Pasca"

var ab, cd float64 = 3.14, 3.15

func main() {
	fmt.Printf("Hello %v your age is %d \n", a, b) // The Printf function are used to format and print text according to a format specifier using verbs.
	fmt.Printf("The data types of a is: %T \n", a) // Printf digunakan untuk print verbs yang berisi %v(default format value) / %T(data type)/ %d(decimal value) , verb %v mencetak value PI yang di ambil dari const PI. variable PI yang berada di line Printf digunakan untuk memasukan value variable PI ke %v. %T digunakan untuk melihat apa tipe data dari variable.
	fmt.Printf("The data types of b is: %T\n", b)

	fmt.Println(a) //the Println function are used to print variable or any value with a spacing and end line at the end of data/variable
	fmt.Println(b)
	
	fmt.Print(a, d) // The Print function are used to print an argument with their defaul value (without spacing, end line etc)
	fmt.Print(b, c) // If the data types is integer or decimal it will automatically add a spacing at the end of each variable as to not make a confusion
	fmt.Print(ab,cd)
}