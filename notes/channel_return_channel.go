package notes

import (
	"fmt"
	//"sync"
	//"time"
)

func generator(s string ) chan int {
    out:=make (chan int)
       go  func (){

              for i:=0 ; i<=5 ;i++{
                  out<-i
                  fmt.Println(i)
              }
              close(out)

        } ()
        
    return out
}
func puller(c chan int  )chan int{
    out:=make(chan int)
    go func (){
           var sum int 
           for n:=range c{
               sum+=n
             }
             out<-sum
             close(out)
    } ()
    return out
}

func dothis(){
    n1:=generator("Genarate1")
    p1:=puller(n1)
    fmt.Println("Sum " , <-p1)

}
