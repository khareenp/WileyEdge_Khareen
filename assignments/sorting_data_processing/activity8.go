//Khareen Proverbs
// Write a program that prompts the user for two different dates, 
// computes the number of days between those dates, and displays the results.

// The order in which the user enters the dates should not affect the results.


package sortingdataprocessing

import(
	"fmt"
	"time"
)


func dateRange(){
	var d time.Duration = 1000000000
	var start,end string
	layout := "2006-01-02"

	fmt.Println("Please enter a start date format: 2006-01-02")
	fmt.Scanf("%s\n",&start)
	fmt.Println("Please enter a end date format: 2006-01-02")
	fmt.Scanf("%s\n",&end)

	

		t1, err := time.Parse(layout, start)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t1)
		t2, err := time.Parse(layout, end)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t2)
	difference := t2.Sub(t1)
	fmt.Printf("difference = %v\n", difference)


}
