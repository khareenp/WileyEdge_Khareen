package notes

import (
	"fmt"
	"time"

)

func loading() {
	const n = 45
    //loading:= spinner()
	fibN := fib(n) // naive and slow
    go spinner()
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

// func spinner(delay time.Duration) {
//     	s:= [8]string{"←", "↖", "↑", "↗", "→", "↘", "↓", "↙"}
//         for i:=0; i < len(s); i++{
//             fmt.Printf("%s",s[i])
//             //fmt.Println("\r")
//             time.Sleep(1 * time.Second)
//         }

// }
func spinner() {
	s := []string{"←", "↖", "↑", "↗", "→", "↘", "↓", "↙"}
	for i := 0; ; i++ {
		fmt.Printf("\r%s", s[i%len(s)-1])//must modulo i or will return value instead of character
		time.Sleep(100 * time.Millisecond)
	}
}

func fib(x int,) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}