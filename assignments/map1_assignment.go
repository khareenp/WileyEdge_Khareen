//Khareen Francis Proverbs
// 1)Create a map where values must be a structure type

package assignments

import (
	"fmt"
)


func MapwithStruct(){
	
	

	//create map where values must be structure type
	//
	type car struct { 
		model string
		brand string 
		color string
	} 

	//make garage (map) 
	garage := make(map[string]*car) 
	 
	//make supra
	car1 := &car{brand: "toyota",model:"supra", color: "red"}	 
 
	//put toyota in garage 
	garage[car1.brand] = car1 
  

	fmt.Println("The first car in the garage is: ",car1.brand,car1.model,car1.color)



}