package assignments

import "fmt"


func increment(arr[]int) {
	//creates array of pointers
	var numbersptr [3] *int
	
	for  i := 0; i < len(numbersptr); i++ {
		// adds address of each value into arr numbersptr
		numbersptr[i] = &arr[i]

		//creates new accessing the values in numbersptr and increment by 1
		newarr := *numbersptr[i]+1
		fmt.Println(numbersptr)
		fmt.Println(newarr)
	  }
}

func Pointerss() {
  // create a simple array
  numbers := []int{100,1000,10000} 

  increment(numbers)
}