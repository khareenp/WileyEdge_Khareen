package map1

import (
	"fmt"
	"sort"
)

func mapfromfn(item map[string]float64) map[string]float64{
    db:=make(map[string]float64)
    for k,v:=range item{
        db[k]=v*2

    }
    return db
}

func Map1(){
	var myGreeting = make(map[string]string)
	//name [" key"] = value ( type specified when declared )
	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."
    myGreeting["Adam"] = "Good Evening."
    myGreeting["Xary"] = "Hello."
	myGreeting["Billy"] = "Bonjour."


	//when iterating map, keys are random,
	// 
	fmt.Println(myGreeting)
    for i,j:=range myGreeting{
        fmt.Println(i,j)
    }

	//short hand creation
    myNums := map[string]string{
		"zero":  "Good morning!",
		"one":   "Bonjour!",
		"two":   "Buenos dias!",
		"three": "Bongiorno!",
	}

	fmt.Println(myGreeting)
	delete(myNums, "two")//delete key value pair 
	


	//SORTING MAP
	//you must first sort by creating a slice
	//and store keys in the slice 
	temp:=map[string]int{"GHI":7,"ABC" : 5, "DEF" : 6,"ZXX": 8 }    
	keys :=make([] string, 0,len(temp))
	for i:= range temp{
    keys=append(keys,i)
				}
	sort.Sort(sort.Reverse(sort.StringSlice(keys)))
	for _,j:=range keys{
    fmt.Println(j,temp[j] )
	}


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
 
	//change original supra (ONLY WORKS if using pointer) 
	garage[car1.brand].color = "green" 
	 
	//also works... 
	//a.color = "green" 


//______________________________
//map using all datatypes
   mapData:=map[string]interface{}{
       "Name" : "noknown", 
       "Age" : 23,
       "Admin" :true,
       "Hobbies":[]string{"IT", "Travel"},
       "Complex" :struct{
           real int
           img int

       }{12,-1},
   }
   fmt.Println(mapData )

   //----------------------------
   //map with key type "string" stores "float"
   p:=make(map[string]float64)
   p["iPhone"]=99256.50
   p["Android"]=5000
   updatadata:= mapfromfn(p)
   for key,val:= range updatadata{
       fmt.Printf( "%s   %v \n" ,key, val)

}

}
 
