package error_handling

import (
   "errors"
   "fmt"
)

type account struct {
   number string
   balance float64
}

func(a *account) withdraw(value float64) (bool,error) {
   if a.balance >= value{
      a.balance = a.balance - value
      return true,nil
   }

   // use the errors package to display a new, custom error message
   return false, errors.New("You cannot withdraw from this account.")
}

func custom() {
   a := account{}
   a.number = "C21345345345355"
   a.balance = 159
   out,err := a.withdraw(200)
   fmt.Println(out)

   // output if error
   if err != nil {
      fmt.Println(err)
      return
   }

   // output if no error
   fmt.Println("The withdrawal occurred succcessfully")
   fmt.Println("Your new balance is", a)
}