//Khareen Proverbs


// Write a function that prompts the user for a year and checks if it is a leap year. 
// The function should return True if the input is a leap year and False otherwise.

// Make sure to test your code with various inputs.

package sortingdataprocessing

import (
	"fmt"
)


func isLeap()bool{
	var year int
	fmt.Println("Please enter a year:")
	fmt.Scanf("%d\n",&year)
	if year%4==0 && year%100!=0 || year%400==0{
        fmt.Println("true")
		return true
	 }else{
         fmt.Println("false")
         return false
}
}

func isleap(){
	isLeap()

}