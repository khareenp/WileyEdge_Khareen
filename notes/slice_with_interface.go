package notes

import "fmt"
func test(){
    i:=[]interface{}{1,2,3}
    j:=[]interface{}{"ABC", "CDE","EFG"}
    

    k:=append(i,j...)
    fmt.Println(k)
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
