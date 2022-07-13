package assignments

//Khareen Proverbs
// Create a program that declares two anonymous functions.

// One function counts down from 100 to 0
// One function counts up from 0 to 100.
// Display each number with a unique identifier for each goroutine.

// Create goroutines from these functions/

// Don't let main return until the goroutines complete.
// Run the programs in parallel.

import (
	"fmt"
	"sync"
)

func range_channel() {
	// Declare a wait group and set the count to two.
	var wg sync.WaitGroup
	wg.Add(2)
    c:=make(chan int)
    d:=make(chan int)

	fmt.Println("Start Goroutines")

	//anonymous function count from 0-100
    go func(){ for i:=0;i<=100 ;i++{
        c<-i
    } 
    close(c)
     }()
     for n:=range c{
         fmt.Printf("First:%d\n",n)
     }

	//anonymous function count from 0-100
    go func(){ for i:=100;i>=0 ;i--{
        d<-i
    } 
    close(d)
     }()
     for n:=range d{
		fmt.Printf("Second:%d\n",n)
	}

}






