package sort

import (
   "fmt"
   "sort"
)

func sortString() {
   // define a slice
   words := []string{"camel", "zebra", "horse", "dog", "elephant", "giraffe"}
   fmt.Println("Original slice:", words)

   
   // sort the values in the slice
   sort.Strings(words)
   fmt.Println("Sorted slice:", words)
   fmt.Println("The values are sorted:", sort.StringsAreSorted(words))//checks if slice is sorted and returns bool
}
