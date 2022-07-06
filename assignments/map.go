package assignments

import "fmt"
func pop(arr []int ,x int)(int ,[]int){
    n:=arr[x]
    return n,append(arr[:x], arr[x+1:]...)//
}
func revese(arr []int)[]int{
   var narry[] int
   for i:=len(arr)-1;i>=0;i--{
       narry=append(narry,arr[i])
   }
   return narry

}


func mapsEven(arr[] int ,f func(int) int ) [] int{
    
    for i,v:=range arr{
        if f(v)%2==0{
            arr[i]=f(v)
        }
        arr[i]=f(v)
    }
    return arr
}
func maps(arr[] int ,f func(int) int ) [] int{
    
    for i,v:=range arr{
        arr[i]=f(v)
    }
    return arr
}

//sums all elements in slice and returns the sum of integers
func mapSum(arr[] int)  int{
	var sum int   
	for i:=0; i < len(arr); i++{
		sum = sum+arr[i]
	}
    return sum
}



func Map(){

arr:=[]int{0,1,2,3,4,5,6,7,8,9,10,11,12}
//fmt.Println(  revese (arr ))
arr= maps(arr,func(r int) int{     
    return r*2 
})
fmt.Println(mapSum(arr))
fmt.Println( arr )



}




