//Khareen Proverbs
// Write a program that adds five seconds to the current time and displays the result.

package sortingdataprocessing
import (
	"fmt"
	"time"
)

func addSec(){
	now:=time.Now()

	fmt.Println(now.Add(time.Second *5))
}