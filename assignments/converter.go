//Khareen Francis Proverbs
//Convert centimeter to inches
 package assignments

 import "fmt"
 
 func Convert(){
	var height_c, height_i, height_f float32
	 height_i= (height_c)/(2.54)
	 height_f= (height_i)/(12)
	 
	fmt.Println("Please enter height in centimeters:")
	fmt.Scanf("%f",&height_c)

	fmt.Println("Height in inches:",height_i)
	fmt.Println("Height in feet:",height_f)
	
	
 }