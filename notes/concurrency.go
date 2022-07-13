package notes

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)


var wg sync.WaitGroup
var Counter int

func concurrency() {
	wg.Add(2)
	go increment("Foo")
	go increment("Bar")
	wg.Wait()
		}
	func increment(s string) {
		rand.Seed(time.Now().UnixNano())
		for i := 0; i <20; i++ {
			x:=Counter
			x++
			time.Sleep(time.Millisecond*1)
			Counter=x
			fmt.Println(s , i, " Counter" , Counter)
		}
		wg.Done()
}
