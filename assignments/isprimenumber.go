//Khareen Proverbs
//Check if any given number is a prime number

package assignments

import (
	"fmt"
)
func CheckPrime(number int) {  
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