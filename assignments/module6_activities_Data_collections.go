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

//--------------------------------------------------------------
//Activity1
func RandomGen(arr int) []int{
	//must use to keep track of machine time and generate new random number
	//or number will repeat
	rand.Seed(time.Now().UnixNano())
	random:=make([]int,arr)//makes slice of size "arr"
	min:= -100//min range
	max:=100//max range
	for i:=0 ; i < arr ; i++{
	//for each position in slice generate a random between max an min value
	random[i] = (rand.Intn(max-min+1)+min)

	}
	return random//return slice
}
//--------------------------------------------------------------
// Activity 2 part 1

func Max(arr []int) int {
	//initialize first value as the max to compare
	max := arr[0]

	//iterate each element in slice
	for i := 1; i < len(arr); i++ {
		//if value in array[0] is smaller
		if max < arr[i] {
			//next value becomes max
			max = arr[i]
		}
	}
    //fmt.Println(max)
	return max
}
//--------------------------------------------------------------
//Activity 2 part 2
//find the index of the max value
func indexOfMax(arr []int)int{
	var indexMax int
	//vallue to store the index number
	for v,i:=range arr{
		indexMax=Max(arr)//index of max value
		//if value "i", is == to the number returned as max 
		//stored in indexMax
		if i == indexMax{
			//display the index number 'v'
			fmt.Println("the index of max number is",v)
		}
	}
	return indexMax
}
//--------------------------------------------------------------
//Activity 2 part 3
func Min(arr []int) int {
	//initialize first value as the minimum to compare
	min:=arr[0]

	//iterate each element in slice
	for i:=1 ; i<len(arr) ; i++{
		//if value in array[0] is larger
		if min > arr[i] {
			//next value becomes minimum
			min=arr[i]
		}
	}
	return min
}

//--------------------------------------------------------------
//Activity 2 part 4
//find the index of the minimum value
func indexOfMin(arr []int)int{
	//vallue to store the index number
	var indexMin int
	for v,i:=range arr{

		indexMin=Min(arr)//index of min value
		//if value "i", is == to the number returned as minimum 
		//stored in indexMin
		if i == indexMin{
			//display the index number 'v'
			fmt.Println("the index of min number is",v)
		}
	}
	//must return to system so it can be used
	return indexMin
}
//--------------------------------------------------------------

//Activity 2 part 6
//Sort an array of int in ascending order and return the new sorted array in a separate array.
func Sortascend(arr []int)[]int{
	//vallue to store the index number
	sorted:=make([]int,len(arr))//slice to store sorted
	length:= len(arr)-1
    
	for i:= 0 ; i < length ; i++{
		for j:= i+1; j <length; j++{
			//tmp:=0
			if arr[i]>arr[j]{
				tmp :=arr[i]
				arr[i]=arr[j]
				arr[j]=tmp
			}
		}
		sorted = append(sorted, arr[i])
	}
    fmt.Println(sorted)
	return sorted
}
//--------------------------------------------------------------
//Activity 2 part 5
//Sort an array of int in descending order and return the new sorted array in a separate array.
func Sortdescend(arr []int)[]int{
	//vallue to store the index number
	sorted:=make([]int,len(arr))//slice to store sorted
	length:= len(arr)-1
    
	for i:= 0 ; i < length ; i++{
		for j:= i+1; j <length; j++{
			//tmp:=0
			if arr[i]<arr[j]{
				tmp :=arr[i]
				arr[i]=arr[j]
				arr[j]=tmp
			}
		}
		sorted = append(sorted, arr[i])
	}
    fmt.Println(sorted)
	return sorted
}
//--------------------------------------------------------------
//Activity 2 part 7
//Compute the mean of an array
func mean(arr []int) int {
    if len(arr) == 0 {
        return 0
    }
    var sum int
    for _, i := range arr {
		//add values of each index and store in sum
        sum += i
    }
    return sum / int(len(arr))
}
//--------------------------------------------------------------
//Activity 2 part 8
//Compute the mean of an array

//--------------------------------------------------------------
//Activity 2 part 9
//identify all positive integers
func positive(arr []int) []int{
	positive_arr:=make([]int,len(arr))//slice to store sorted
	for i:=0; i < len(arr); i++{
		if arr[i]>0{
			positive_arr= append(positive_arr, arr[i])
		}
	}
	return positive_arr
}

//--------------------------------------------------------------
//Activity 2 part 10
//identify all negative integers
func negative(arr []int) []int{
	positive_arr:=make([]int,len(arr))//slice to store sorted
	for i:=0; i < len(arr); i++{
		if arr[i]<0{
			positive_arr= append(positive_arr, arr[i])
		}
	}
	return positive_arr
}

//--------------------------------------------------------------



func Start(){
	
	slice_100:=RandomGen(100)
  
    fmt.Println(slice_100)
	fmt.Printf("The largest number is %d\n",Max(slice_100))
    fmt.Printf("The smallest number is %d\n",Min(slice_100))
    indexOfMax(slice_100)
    indexOfMin(slice_100)
	fmt.Println("\nSorted in Descending order:")
    Sortdescend(slice_100)
	fmt.Println("\nSorted in Ascending order: ")
    Sortascend(slice_100)
	fmt.Println("\nThe mean of slice_100 is: ",mean(slice_100))
	fmt.Println("\nThe positive numbers are: ",positive(slice_100))
	fmt.Println("\nThe negative numbers are: ",negative(slice_100))

   
}
