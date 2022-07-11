package error_handling

import (
   "fmt"
   "strconv"
)

func example1() {
   var s string = "10x"

   // the ParseInt function returns the parsed integer or
   // the error if the conversion failed
   a, error := strconv.ParseInt(s,10,8)
   fmt.Println(a)
   fmt.Println(error)
}