//Khareen Proverbs
// Reverse char
package assignments
import (
	"fmt"
	"unicode"
)

func ReverseChar() {
    value:= "Hello World"

   new_val:=""
   for pos,val:= range value{

      if unicode.IsSpace(val){
         for i:=pos-1; i >=0; i--{
            new_val = new_val+ string(value[i])
         }
         new_val = new_val+ " "
         for i:=len(value)-1; i> pos; i--{
            new_val += string(value[i])
         }
      }
   }
   fmt.Println(new_val)
}