//Khareen Francis Proverbs
//Activity 10
/**/
//Write a program that takes as input a number and produces a 
//triangle as shown below:



package assignments
import ("fmt")

func Triangle(){
	var rows int

	fmt.Println("Please enter number of rows:")
	fmt.Scanf("%d",&rows)

	for i:=1; i<=rows;i++{
        for j:=1;j<=i;j++{
            fmt.Printf("%d",j)
        }
        fmt.Println(" ")
    }
}