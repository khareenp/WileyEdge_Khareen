//Khareen Proverbs

// Write a program that prints the dates for yesterday, today, and tomorrow.

package sortingdataprocessing

import(
	"fmt"
	"time"
)

func printDates(){
	now:= time.Now()
	yesterday:=now.Add(-time.Hour *24)
	tomorrow:=now.Add(time.Hour *24)

	fmt.Println("Today's date is",(now).Format("Jan 2, 2006"))
	
	fmt.Println("Yesterday's date is:",(yesterday).Format("Jan 2, 2006"))
	fmt.Println("Tomorrow's date is:",(tomorrow).Format("Jan 2, 2006"))
}

func getDates(){
    printDates()
}