// Khareen Francis Proverbs
// Activity 21

// An Armstrong number is a number where the sum of the cubes 
// of each digit is equal to the number itself. 
// For example, 153 is an Armstrong number:
// Write a program that identifies and prints out all 
// Armstrong numbers between 1 and 500.

package assignments

import (
	"fmt"
	//"strconv"
)

func Armstrong(){
	var orig,number int
	result:=0
	orig = number

	for number !=0{
		remainder:= number%10
		result = result + (remainder*remainder*remainder)
		number = number/10

	}
	//number is Armstrong
	if orig == result{
		res:= orig==result
		fmt.Println(bool(res))
	}

	

}