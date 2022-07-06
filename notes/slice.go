//Khareen Francis-Proverbs

package notes

import (
	"fmt"
	//"reflect"
 )
 
 func slice() {
//    // define an array with 7 elements
//    numbers := [7]int{0,1,2,5,798,43,78}
//    fmt.Println("array value:", numbers)
//    fmt.Println("array type:", reflect.TypeOf(numbers))
 
//    // define a slice s based on the numbers array
//    s := numbers[0:4]
//    fmt.Println("slice value:", s)
//    fmt.Println("slice type:", reflect.TypeOf(s))

//    fmt.Println("Length of slice:", len(s))//shows length based on elements

//    fmt.Println("Slice capacity:", cap(s)) //shows capacity based on original array

//create double dimension slice
 	var records [][] string

 //"make" used to create slice single dimension 
 //of type string and length 1

	students1:=make([]string,1)
	students1[0]="ABC"
	students1[1]="XYZ"
	students1[2]="AXZ"

	student2:=make([]string,3)
	student2[0]="123"
	student2[1]="456"
	student2[2]="789"


	records=append(records,student2)
	fmt.Println(records)
 }





 /*
The output:
array value: [0 1 2 5 798 43 78]
array type: [7]int
slice value: [0 1 2 5]
slice type: []int
*/