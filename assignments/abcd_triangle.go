// Khareen Francis Proverbs
// Alphabet triangle

package assignments

import "fmt"

func Abcd() {
    asc:=65

	for i := 1; i < 5; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf(("%c"), asc+j)
		}
		fmt.Println()
	}
}