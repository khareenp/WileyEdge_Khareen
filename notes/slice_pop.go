package notes

import "fmt"
func pop(arr []int ,x int)(int ,[]int){
    n:=arr[x]
    return n,append(arr[:x], arr[x+1:]...)//
}
func revese(arr []int)[]int{
   var narry[] int
   for i:=len(arr)-1 ; i >=0 ;i--{
       narry=append(narry,arr[i])
   }
   return narry

}

//takes slice and function as parameter
//will filte reach element of slice
// func maps(arr[] int ,f func(int) int ) [] int{
    
//     for i,v:=range arr{
//         arr[i]=f(v)
//     }
//     return arr
// }

//maps for even numbers only
func maps(arr[] int ,f func(int) int ) [] int{
    
    for i,v:=range arr{
        if f(v)%2==0{
            arr[i]=f(v)
        }
        arr[i]=f(v)
    }
    return arr
}

func run(){

arr:=[]int{0,1,2,3,4,5,6,7,8,9}
//fmt.Println(  revese (arr ))
arr= maps(arr,func(r int) int{     
    return r*2 
})

fmt.Println(  arr )


//a,arr:=pop(arr,6)// index
//fmt.Println(a,arr)


  /*  i:=[]interface{}{1,2,3}
    j:=[]interface{}{"ABC", "CDE","EFG"}
    k:=append(i,j...)
    fmt.Println(k)
    */
}
/*
func main() {
    matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

     for i,j:= range matrix{
        fmt.Println( i, "  ", j[0])
    }
    //you can change the position of slice index 
matrix[0],matrix[2]=matrix[2],matrix[0]
    fmt.Println(matrix)
    for i,j:= range matrix{
        fmt.Println( i, "  ", j)
    }
}

*/

