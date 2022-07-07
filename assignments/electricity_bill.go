//Khareen Francis Proverbs
//switch assignment

/*
Electricity bill from 0 to 100 unit rate $5 per unit 
from 101 to 200 unit $7 per unit for 201 to 350 unit $10 
per unit calculate total bill

*/

package assignments

import "fmt"



func Electricity(){
	var units,total float32

	fmt.Println("Please enter the number of units in electrical bill:")
	fmt.Scanf("%f",&units)


	if (units >=0 && units <=100){
        total += units*5
        fmt.Println("Total is",total)

    } else if (units >=101 && units <=200){
        base := (100*5)
        difference:= units-100
        new_rate := difference *7
        total:= float32(base) + new_rate
        fmt.Println("Total is",total)

    }  else if (units >=201 && units <=350){
        base := (100*5)
        difference:= units-100
        new_rate := difference *10
        total:= float32(base) + new_rate
        fmt.Println("Total is",total)
    } 
    

}