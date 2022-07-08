//Khareen Francis Proverbs
//Module 6 Activities: Data Collections

//Activity 1
//Create a function that takes as input an int n and returns 
//an array of length n with random integers between -100 and 100.
//_____________________________________________________________________
// Activity 2
// Using the previous function, create an array of 100 numbers.

// Implement a function for each of the following bullets. 
// Each function will take as an input an array of 100 integers. 
// Do not use the sort package or any 
// built-in function (like min, max, etc.).

// *Compute the max of an array of int
// *Compute the index of the max of an array of int
// *Compute the min of an array of int
// *Compute the index of the min of an array of int
// *Sort an array of int in descending order and return the new sorted array in a separate array.
// *Sort an array of int in ascending order and return the new sorted array in a separate array.
// *Compute the mean of an array
// *Compute the median of an array
// *Identify all positive numbers in the array and return the numbers in a slice
// *Identify all negative numbers in the array and return the numbers in a slice
// *Compute the longest sequence of sorted numbers (in descending or ascending order) in the array and return in a new array
		// Example: input: [1 45 67 87 6 57 0]
		// Output: [1 45 67 87]
// *Remove duplicates from an array of ints and return the unique elements in a slice


package assignments

import (
	"fmt"
	"math/rand"
	"time"

)


func RandomGen(n int) []int{
	//must use to keep track of machine time and generate new random number
	//or number will repeat
	rand.Seed(time.Now().UnixNano())
	random:=make([]int,n)//makes slice of size "n"
	min:= -100//min range
	max:=100//max range
	for i:=0 ; i < n ; i++{
	//for each position in slice generate a random between max an min value
	random[i] = (rand.Intn(max-min+1)+min)

	}
	return random//return slice
}


func Max(n []int) int {
	//initialize first value as the max to compare
	max := n[0]

	//iterate each element in slice
	for i := 1; i < len(n); i++ {
		//if value in array[0] is smaller
		if max < n[i] {
			//next value becomes max
			max = n[i]
		}
	}
    //fmt.Println(max)
	return max
}

func Min(n []int) int {
	//initialize first value as the minimum to compare
	min:=n[0]

	//iterate each element in slice
	for i:=1 ; i<len(n) ; i++{
		//if value in array[0] is larger
		if min > n[i] {
			//next value becomes minimum
			min=n[i]
		}
	}
	return min
}

func indexOfMax(n []int)int{
	var indexval int
	for v,i:=range n{
		indexval=Max(n)
		if i == indexval{
			fmt.Println("the index of max number is",v)
		}
	}
	return indexval
}
func indexOfMin(n []int)int{
	var indexval int
	for v,i:=range n{
		indexval=Min(n)
		if i == indexval{
			fmt.Println("the index of max number is",v)
		}
	}
	return indexval
}



func Start(){
	
	slice_100:=RandomGen(100)
  
    fmt.Println(slice_100)
    fmt.Println(Max(slice_100))
    fmt.Println(Min(slice_100))
    indexOfMax(slice_100)
    indexOfMin(slice_100)

	
   
}
