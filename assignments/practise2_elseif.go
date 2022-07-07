//Khareen Francis Proverbs
//Practise 2 Esle if

package assignments

import (
	"fmt"
)

/*
A grade calculator that outputs a letter
grade based on a given numeric grade
*/

func ElseIfPractise() {
   
  
    grade:=55

   fmt.Println("Please enter grade: ")
   fmt.Scanln(&grade)

   if grade >= 90 {
	fmt.Println("You received an A")
} else if (grade >= 80 && grade <90) {
	fmt.Println("You received a B")
} else if (grade >= 70 && grade <80) {
	fmt.Println("you received a C")
} else if (grade >= 65) && grade <70 {
	fmt.Println("You received a D")
} else {
	fmt.Println("Sorry, you have failed")
}
   
    
}