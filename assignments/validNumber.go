// Use regx to find valid number
package assignments

import (
	"fmt"
	"regexp"
)

func validnum() {
	num1 := "1(234)5678901x1234"
	num2 := "(+351) 246 55 29"
	num3 := "11122233451"
	num4 := "asdf"
	
	re := regexp.MustCompile(`1?\W*([2-9][0-8][0-9])\W*([2-9][0-9]{2})\W*([0-9]{4})(\se?x?t?(\d*))?`)

	fmt.Printf("Pattern: %v\n", re.String()) // print pattern
	fmt.Printf("\nPhone: %v\t:%v\n", num1, re.MatchString(num1))
	fmt.Printf("Phone: %v\t:%v\n", num2, re.MatchString(num2))
	fmt.Printf("Phone: %v\t\t:%v\n", num3, re.MatchString(num3))
	fmt.Printf("Phone: %v\t\t:%v\n", num4, re.MatchString(num4))
	
}