// For instance, imagine that we want to sort a set of words based on
// the number of characters in each word, rather than alphabetically.
// The Sort package include the Sort interface, which we can implement
// with our own logic to create custom sorting algorithms.

// The sort interface includes three methods that we need to implement:
//  Len, Swap, and Less.

// Len:
// The Len function returns the length of the data type
// in context of sorting. In this case, we want to sort words based
// on their length so the Len function will simply return
// the length of the input words.

// Swap:
// The sort package uses the Swap function internally to swap
// items in a slice or array during sorting.
// The function takes as input two indexes and it swaps the values
// in those indexes.

// Less:
// The Less function provides the logic for comparing two items in a
// slice or array. Because we want to compare based on length,
// this function will take the indexes of the two items we want to
// compare in the slice or array, and it will return true if the
// length of the first word (stored in the first index) is higher
// than the length of the second word (stored in the second index).
// The following code implements the three methods of the sort
// package to sort the string values based on the length of the string.

package sort

import (
	"fmt"
	"sort"
)
//In order to implement the interface, 
//we need our own custom type,
// create an alias type
type mytype []string

// implement the Len method
func (s mytype) Len() int {
   return len(s)
}

 // implement the Swap method
func (s mytype) Swap(i, j int) {
   s[i], s[j] = s[j], s[i]
}

// implement the Less method
func (s mytype) Less(i, j int) bool {
   return len(s[i]) < len(s[j])
}

func sortStringLen() {
   // create a slice of strings
   fruits := []string{"pear", "passionfruit", "mango", "banana", "fig"}
   fmt.Println("Original slice:", fruits)

   // create a mytype variable
   myfruits := mytype(fruits)

   sort.Sort(myfruits)
   fmt.Println("Sorted by length:", myfruits) 
}