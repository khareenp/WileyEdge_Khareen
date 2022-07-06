package notes

import (
   "fmt"
)

func iterate() {
   numbers := [7]int{0,1,2,5,798,43,78}
   fmt.Println(numbers)

   s := numbers[0:4] 
   for i := 0; i < len(s); i++{
      fmt.Println("Element", i, "is", s[i])
   }
}
