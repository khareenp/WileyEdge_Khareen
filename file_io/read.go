package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
   // use the ReadFile from ioutil package to read the entire file in memory
   data, err := ioutil.ReadFile("/Users/kp/Desktop/wiley edge/Submit/file_io/flatland01.txt")

   // feedback message in case of error
   fmt.Println(err)

   // convert the file contents to a string and display it
   fmt.Print(string(data))
}
