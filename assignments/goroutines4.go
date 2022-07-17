//Khareen Proverbs
//Friday July 15
//go routines and channels

// 4)Create Two go routine where routine 1 generate random number and
// append them in  slice where another goroutine sort the slice at the same time .
// print the slice after every append and sorted array at the same time  side by side

package assignments

import (
	"fmt"
	"math/rand"
	"sort"
)

func gorandd(max int,min int, arr []int, ch chan<- []int) {
	minn := min
    maxx := max
	newarr:= make([]int,len(arr)-1)
	for _, i := range arr {
		i= (rand.Intn(maxx - minn) + minn)
		newarr= append(newarr,arr[i])
		fmt.Println(newarr)
		
	}
	ch <- newarr
	close(ch)
}

func sortchann(arr []int, ch chan<- []int) {

		s:= sort.IntSlice(arr)
		fmt.Println(s)
		ch <- s
		
	close(ch)
}

func gorand(){

	max:=50
	min:=1
	var arr []int
	ch1:= make(chan []int)

	gorandd(max,min,arr,ch1)
	sortchann(arr,ch1)

}