//Khareen Francis-Proverbs
//July 1st 2022 

package assignments
import (
	"fmt"
	"math"
)

// 1.sum is a function which takes a slice of numbers and adds them together. 
// What would its function signature look like in Go?

func sum(slice[] int) int{
	sum:= 0
	for _, i := range slice {
		sum= sum+i		
	}
	return sum
}



// 2)Write a function which takes an integer and halves it and returns true if it was even or false
//  if it was odd. For example half(1) should return (0, false) and half(2) should return (1, true).

func half(num float64){
	half:= num/2
	res := math.Mod(half, 2)
	//half_num:= half %2 
	
	if res==0{
		fmt.Println(res,"true")
	} else{
		fmt.Println(res,"false")
	
	}

}

// 3). Write a function with one variadic(variable numbers) parameter that finds
//  the greatest number in a list of numbers.

func largest(numbers ... int) int {
	max_num := 0
	for _, i := range numbers {
		if i > max_num {
			max_num = i
		}
	}
	return max_num
}

// 4)Using makeEvenGenerator as an example, write a makeOddGenerator 
// function that generates odd numbers.

func makeOddGenerator() func() uint {
    i := uint(1)
    return func() (num uint) {
        i += 2 
        return i
    }
}

func evenGenerator() func() int {
	i := int(0)
	return func() int {
		i += 2
		return i
	}
}



// 5)The Fibonacci sequence is defined as: 
// 0, fib(0) = fib(1) = 1, fib(n) = fib(n-1) + fib(n-2). 
// Write a recursive function which can find fib(n).

func fib(n uint) uint {
    if n == 0 {
        return 0
    } else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func Weekly() {

	
}