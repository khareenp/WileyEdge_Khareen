//Khareen Proverbs
//Exponential or sin series
//activity 8 modify

/********************************
************SIN Series***********
Series: sin(x) = x - (x^3/3!) + (x^5/5!) + .....
********************************/

package assignments

import (
	"fmt"
)

func Sin(){

	
	//Declare variables
	var i,n int
	var x, term1, term2, R, sum float32

	//Prompt user for input
	fmt.Println("Please enter the value of x:")
	fmt.Scanf("%f",&x)

	//Prompt user for input
	fmt.Println("Please enter the number of values to be summed:")
	fmt.Scanf("%d",&n)

	//Initialize First Term
    term1 = x

    //Make sum equal to the first term
    sum = term1

	fmt.Printf("n\ttn\t\tSn\n_________________________________");

	//condition to loop
	for i=1;i<n;i++{
        //Find the ratio of the second term to the first term using already known relation
        R=- float32((x*x))/float32((2*i+1))/float32((2*i))

        //Calculate the second term
        term2 = R*term1

        //find the new sum
        sum=sum+term2
        term1=term2
        fmt.Printf("\n%d\t%f\t%f\n",i+1,term2,sum)
    }
    fmt.Printf("\nThe sum is: %f",sum)
}




