
//Khareen Proverbs
// Write a program that prints the date for the next five days, starting from today.

package sortingdataprocessing

import(
	"fmt"
	"time"

)

func nextdays(){

		var dates [5]string
	
		now:= time.Now()
		tomorrow:=now.Add(time.Hour *24)
	
		fmt.Println("Today's date is",(now).Format("Jan 2, 2006"))
		
		fmt.Println("Tomorrow's date is:",(tomorrow).Format("Jan 2, 2006"))

		for i:=1; i < len(dates);i++{
			dates[i]+= tomorrow.String()
			fmt.Println(dates)
		}

}