//Khareen Proverbs
//Activity 11
// Write a program that calculates the greatest common divisor between
// two numbers.
// For example, given the numbers 18 and 27, the greatest common 
// divisor is 9.



package assignments
import ("fmt")

func Gcd(){
	var num1,num2,gcd int

	//prompt user for input
	fmt.Println("Please enter number 1:")
	fmt.Scanf("%d",&num1)
	fmt.Println("Please enter number 2:")
	fmt.Scanf("%d",&num2)

	//while loop will stop when either number is 0
	for num1!=0 && num2!=0{
		num3:= num2
		num2 = num1%num2
		num1= num3
		gcd=num1+num2 //either value may be 0, add values to get non-zero value
	}
	fmt.Printf("The GCD of values is: %d",gcd) //print result
}