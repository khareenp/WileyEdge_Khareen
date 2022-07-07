//Khareen Francis Proverbs
// Map sorting

// 3) Create a map ex
//string key value int
// map[string]int{"orange": 5, "apple": 7,	"mango": 3, "strawberry": 9} ,
// sort the map based on key length in ascending order
//based on length of key from small to larger


// 2) Take 20 (any count) from console between 1 -100 ) 
// then print the summary like no between 1- 10 Count-5     
// 11 -20 count -7    21-30 count -10  etc

package assignments

import (
	"fmt"
	"sort"
)

func MapSort() {
	//create map with string key and int values
	newMap := map[string]int{"orange": 5, "apple": 7, "mango": 3, "strawberry": 9}

	//create slice to store keys with capacity(maplen- will be a number based on elements in the map)
	keys := make([]string, 0)
	//keyLen:=make([] string, 0,len(keys))

	for i := range newMap {
		keys = append(keys, i)
	}

	sort.Slice(keys, func(i, j int) bool {
		return len(keys[i]) < len(keys[j])
	})

	for _, k := range keys {
		fmt.Printf("k = %v, v = %v\n", k, newMap[k])
	}
}