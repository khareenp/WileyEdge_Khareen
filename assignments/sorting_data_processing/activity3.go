//Khareen Proverbs

// Write a function that prompts the user for a date
// (or uses the current date if the user does not input a date) and
// subtracts five days from that date. The function should return the new date.

// Display the result to the user.

package sortingdataprocessing

import (
	"fmt"
	"time"
)

func acceptdate(){
	now := time.Now()

	var input string
	fmt.Println("Please enter a date in the format:Jan 2, 2006")
	fmt.Scanf("%d\n",input)

	if input ==""{
		input2 := (now.Add(-time.Hour *120).Format("Jan 2, 2006"))
		
		fmt.Println(input2)
	} else{
		// Parse the date string into Go's time object
		// The 1st param specifies the format, 2nd is our date string
		myDate, err := time.Parse("2006-01-02 15:04", input)
	if err != nil {
		panic(err)
	}
	// Format uses the same formatting style as parse, or we can use a pre-made constant
	fmt.Println("My Date Reformatted:\t", myDate.Format("Jan 2, 2006"))

	}
}