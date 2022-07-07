// Khareen Francis Proverbs
//Twin Primes
// Khareen Francis Proverbs

package assignments

import "fmt"

func TwinPrime() {
	for i := 0; i <= 6; i++ {
		for j := 1; j <= 7-i; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}
}
	
	