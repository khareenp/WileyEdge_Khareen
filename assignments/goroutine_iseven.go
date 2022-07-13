//Khareen Proverbs
//use 2 goroutines to check if a number is odd or even.

package assignments

import "fmt"

import "time"
import "os"


var (
	printNow chan bool
    i int
)

func main() {
	printNow = make(chan bool)
   
    
	go printer()
	go sender()
   	for {
	}
}
func printer() {
	for {
		if _, ok :=<-printNow ;ok{
            switch{
                case i %2== 0:
                fmt.Println("is even")
                case i %2==1:
                fmt.Println("is odd")
            }
			//fmt.Println("Recieved !" , i) 
		}else{
             return
             os.Exit(0)
        }
        
	}
}
func sender() {
	for {
		for i = 0; i < 10; i++ {
            fmt.Println("The number" ,i)
            
			printNow <- true
            
			time.Sleep(1 * time.Millisecond)
		}
        close(printNow)
        return
	}
}
