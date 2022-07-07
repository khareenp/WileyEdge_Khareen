// Khareen Francis Proverbs
//Twin Primes
// Khareen Francis Proverbs

package assignments

import "fmt"

func TwinPrime() {

	var num int
	fmt.Println("Please enter positive integer that will determine the max value to search:")
	fmt.Scanf("%d",&num)


	for i := 0; i <= num; i++ {
		for j := 1; j <= num-i; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}
}
	
	