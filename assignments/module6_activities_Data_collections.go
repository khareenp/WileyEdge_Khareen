//Khareen Francis Proverbs
//Module 6 Activities: Data Collections

//Activity 1
//Create a function that takes as input an int n and returns 
//an array of length n with random integers between -100 and 100.


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

func Random(){
	slice:=RandomGen(20)
    fmt.Println(slice) 
}
