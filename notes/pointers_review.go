package notes

import "fmt"

func changebyVal(x int ){
    x=x+10
    fmt.Println(x ," inside  " , &x)// 20

}
func changebyAddress(x *int ){
     *x=*x+10
    fmt.Println(*x ," inside  " , x)
}

func PointerReview(){
a:=10
//changebyVal(a)
//fmt.Println(a, "  " ,&a)
changebyAddress(&a)
fmt.Println(a, "out side  " ,&a)
}