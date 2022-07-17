//Khareen Proverbs
//Friday Jun 15
//go routines orig


// 2)Take a Paragraphs of text Max word  up to 200 words in a Slice or String
//  and then Pass each  Sentence(up to full stop) to another go routine/routines   
//  1)print it in reverse order(1 goroutine) 
//  2 )count its words(2 goroutine)
package assignments

import (
	"fmt"
	"strings"
)
//
func reverseCH(str []string, ch chan<- string) {
	for _, str := range str {
		runes := []rune(str)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		ch <- string(runes)
	}
	close(ch)
}

func countWordsCH(str []string, ch chan<- int) {
	for _, str := range str {
		ch <- len(strings.Split(str, " "))
	}
	close(ch)
}

func Driver2() {
	st :=  "If you're looking for random paragraphs, you've come to the right place. When a random word or a random sentence isn't quite enough, the next logical step is to find a random paragraph. We created the Random Paragraph Generator with you in mind. The process is quite simple. Choose the number of random paragraphs you'd like to see and click the button. Your chosen number of paragraphs will instantly appear."


	strs := strings.Split(st, ".")

	ch1 := make(chan string)
	ch2 := make(chan int)
	//first go routine to reverse in channel 1
	go reverseCH(strs, ch1)
	//second goroutine to count in channel 2
	go countWordsCH(strs, ch2)

	for str := range ch1 {
		fmt.Println(str)
	}
	for i := range ch2 {
		fmt.Println(i)
	}
}