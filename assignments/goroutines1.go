//Khareen Proverbs
//Friday Jun 15 activity

// 1)Create a slice of 20 of type int and take 20 number 
// create 4 goroutines to take average of every 5 elements ex:
// [1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20] ,
// so 1 goroutine should give average of first 5 elements and next for another 5 so on
// 15,16,17,18,19,20] ,so 1 goroutine


package assignments
import (
	"fmt"
)


func Sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // send sum to c
}

func Average(a []int, d chan int){
	//size of array
	length:= len(a)
	sum := 0
	for _, v := range a {
		sum+= a[v]
	}
	average:= (float64(sum)/float64(length))

	d <- int(average) // send sum to d
}

func Driver() {
	a := [20]int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}

	d:= make(chan int)
	c := make(chan int)
	go sum(a[:len(a)/4], c)
	go sum(a[len(a)/2:], c)

	go Average(a[0:5],d)
	go Average(a[5:10], d)
	go Average(a[10:15], d)
	go Average(a[15:], d)
	x, y := <-c, <-c // receive from c
	e,f,g,h:= <-d, <-d,<-d,<-d

	fmt.Println(x, y)
	fmt.Println(e,f,g,h)
}
