package main

import "fmt"

func main() {

	/*SIGNED INTEGER

		SIGNED INTEGER HAS 5 TYPES, and we are able to store a positive value or negative value.
		the default value of an int is 0
		We have to optimize the integer variable declaration based on what we need, if we need a large number for a bank project we have to use the int64
		the default is int, if we did not specify the signed types it will decide based on OS 32bit or 64bit
		int8 range from -128 to 127
		int16 range from -32768 to 32767
		int32 range from -2147483648 to 2147483647
		int64 range from -9223372036854775808 to 9223372036854775807.
	*/
	var myAge int = 25
	var myCar int8 = 127
	var myTuition int16 = 32767
	var myCardMoney int32 = 2147483647
	var myTotalMoney int64 = 9223372036854775807

	/*UNSIGNED INTEGER
		UNSIGNED INTEGER ALSO HAS 5 TYPES,but we are only able to store a positive value
		the default value of an uint is 0
		the default is uint, if we did not specify the unsigned types it will decide based on OS 32bit or 64bit
		uint8 range from 0 to 255
		uint16 range from 0 to 65535
		uint32 range from 0 to 4294967295
		uint64 range from 0 to 18446744073709551615.
		as you can see it doubles the range values
	*/

	var myUAge uint = 25
	var myUCar uint8 = 127
	var myUTuition uint16 = 32767
	var myUCardMoney uint32 = 2147483647
	var myUTotalMoney uint64 = 9223372036854775807
	fmt.Println(myAge, myUAge)
	fmt.Println(myCar, myUCar)
	fmt.Println(myTuition, myUTuition)
	fmt.Println(myCardMoney, myUCardMoney)
	fmt.Println(myTotalMoney, myUTotalMoney)

	/*Boolean
		boolean consist of True and False
		the default value of boolean is false
	*/

	var isLegal bool = true
	fmt.Println(isLegal)

	/* Float
		there are 2 types of float, float32 and float64
	*/

	var myPi float32 = 3.14
	var myNotPi float64 = 4.14
	fmt.Println(myPi, myNotPi)

	/* String
		String are used to hold a text value
		The default value of string is empty string ""
		string value needs to be in double quotation ""
	*/

	var myString string = "Hello"
	fmt.Println(myString)
}