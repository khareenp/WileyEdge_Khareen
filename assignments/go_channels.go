//Khareen Proverbs
//1 go routine to generate number
// next go routine to print number if prime and where number found
// use channels to communitacte

package assignments

import (
	"fmt"
	//"sync"
	//"time"
)
func findprime(){
    c:=make( chan int )
    go func(){
        for i:=0;i<=9;i++{
            c<-i
        }   
  close(c)
    }()  // Ann function is over 
    
        for ii:= range c{
              fmt.Println(ii)
        }
       // reading 

}
