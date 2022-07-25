//Khareen Proverbs
// Write a program to display the following:

// Current date and time *
// Current year *
// Current month *
// Week number of the year
// Weekday of the week *
// Day of the year
// Day of the month*
// Day of the week *

package sortingdataprocessing

import (
	"fmt"
	"time"
)

func main(){
	now:= time.Now() //current ytd
	yearday:=now.YearDay()
	year:= now.Year()// current year
	month:= now.Month()//current month
	day:=now.Day() //day of the month
	weekday:= now.Weekday() //day of the week

	fmt.Println("The current datetime is:", now)
	fmt.Println("The current year is:", year)
	fmt.Println("The current day of the year is:", yearday)
	fmt.Println("The current month is:", month)
	fmt.Println("The current day of the month is:",day) 
	fmt.Println("The current weekday is:", weekday)
	fmt.Println("")
	

	dateandtime()
}

func dateandtime(){
	now := time.Now()
	fmt.Println("Time: ", now.Format("15:04:05"))
    fmt.Println("Date:", now.Format("Jan 2, 2006"))
    fmt.Println("Timestamp:", now.Format(time.Stamp))
    fmt.Println("ANSIC:", now.Format(time.ANSIC))
    fmt.Println("UnixDate:", now.Format(time.UnixDate))
    fmt.Println("Kitchen:", now.Format(time.Kitchen))

}