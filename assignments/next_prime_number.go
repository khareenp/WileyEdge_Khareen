//Khareen Francis Proverbs
// Generate the next prime number after "n"
package assignments

import (
   "fmt"
   //"math"
)

func isPrime(n int) bool{
	//corner case
	if n <= 1{
		return	false
	}
	if n<=3 {
			return true
		}

	if n%2 == 0 || n %3 == 0{
		return false
	}
	for i:= 5; i*i <= n; i = i+6{
		if n%i ==0|| n%(i+2)==0{
			return false
		}

	}
	return true
}

//function to return smallest
// prime number greater than "n"
func nextPrime(num int) int{
	if num <=1{
		return 2
	}
	prime:= num
	var found bool = false

	//loop until isPRime returns true
	// for a number greater than num

	for !found{
		prime++

		if isPrime(prime){
			found =true
		}
		 
	}
	return prime
}

func NextPrime() {

	var num int
	start:
	fmt.Println("Please enter a positive integer and we will determine if the next prime number:")
    fmt.Scanf("%d",&num)

	if num <=0 {
		goto	MESSAGE1
	}else{
		goto commence
	}

	MESSAGE1: fmt.Println("Sorry you entered an incorrect value. Please try again.")
	goto start

	commence:
	fmt.Println(nextPrime(num))
   
}