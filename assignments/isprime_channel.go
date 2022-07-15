//Khareen Proverbs
//Convert puller functin to prime number checker
package assignments

import (
	"fmt"
	"math"
	//"sync"
	//"time"
)

func generator(s string ) chan int {
    out:=make (chan int)
       go  func (){

              for i:=0 ; i<=5 ;i++{
                  out<-i// write to channel
                  fmt.Println(i)
              }
              close(out)
        } ()
        
    return out
}
func puller(c chan int  )chan int{
    out:=make(chan int)
    go func (){
            
           for n:=range c{
			var prime int
			prime=CheckPrime(n)
               
             }
             out<-prime
             close(out)
    } ()
    return out
}

func CheckPrimenum(number int) {  
	isPrime := true  
	if number == 0 || number == 1 {  
	 fmt.Printf(" %d is not a  prime number\n", number)  
	} else {  
	 for i := 2; i <= number/2; i++ {  
	  if number%i == 0 {  
	   fmt.Printf(" %d is not a  prime number\n", number)  
	   isPrime = false  
	   break  
	  }  
	 }  
	 if isPrime == true {  
	  fmt.Printf(" %d is a prime number\n", number)  
	 }  
	}  
   }  

func doothis(){
    n1:=generator("Genarate1")
    p1:=puller(n1)
    fmt.Println("Sum " , <-p1)

}
