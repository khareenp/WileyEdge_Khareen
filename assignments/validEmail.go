package assignments

import (
	"fmt"
	"regexp"
)

func validemail() {
	e := "test@tester.com"
    b := "test123.com"

 emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	//re := regexp.MustCompile(`1?\W*([2-9][0-8][0-9])\W*([2-9][0-9]{2})\W*([0-9]{4})(\se?x?t?(\d*))?`)

	fmt.Printf("Pattern: %v\n", emailRegex.String()) // print pattern
    fmt.Printf("\n email: %v\t:%v\n", e, emailRegex.MatchString(e))
    fmt.Printf("\n email: %v\t:%v\n", b, emailRegex.MatchString(b))

}