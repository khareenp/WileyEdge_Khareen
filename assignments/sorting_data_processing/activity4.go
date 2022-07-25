//Khareen Proverbs
// Write a program that converts a Unix timestamp string to a human-readable date.

package sortingdataprocessing

import(
	"fmt"
	"time"
)

func convert(){
	unixTime := time.Now().Unix()

	fmt.Println(unixTime)

    

    //Unix Timestamp to time.Time
    convert:= time.Unix(unixTime, 0)
    fmt.Printf("time.Time: %s\n", convert)
}
