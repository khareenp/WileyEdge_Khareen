//Khareen francis Proverbs
// Activity 6
// Write a program that asks the user for a positive integer and then 
// displays the multiplication table of that input number from 1 
// through the integer squared.

package assignments

import(
	"fmt"
)
func Multiply() {
	var num int
	
	start: //label for goto on line 28
	fmt.Println("Please enter a positive integer: ")
	fmt.Scan(&num)
	if num >= 0{
		for i := 1; i <= num*num; i++ {
			fmt.Printf("%v * %v = %v\n", num, i, num*i)
		}
	} else{
		goto MESSAGE1
	}

	MESSAGE1: 
	fmt.Println("Please try again that value is incorrect")
	goto start //will go to this label and execute program
	
}