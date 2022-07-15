//Khareen Proverbs
//Friday Jun 15
//go routines


// 2)Take a Paragraphs of text Max word  up to 200 words in a Slice or String
//  and then Pass each  Sentence(up to full stop) to another go routine/routines   
//  1)print it in reverse order(1 goroutine) 
//  2 )count its words(2 goroutine)

package assignments

import ("fmt"
     "strings"
     "time"
)


//takes a string and a channel of array as parameters
func Send(sample string, ch chan []string){
	res:= strings.Split(sample, ".")
	ch <- res // value received
}

func receive( sample string, ch chan []string){
	time.Sleep(time.Second *2)
	res:= <- ch
	for i:= 0; i < len(res); i++{
		fmt.Println(res[i])
	}
	fmt.Println(res)
}


func Driver2(){
	st := "If you're looking for random paragraphs, you've come to the right place. When a random word or a random sentence isn't quite enough, the next logical step is to find a random paragraph. We created the Random Paragraph Generator with you in mind. The process is quite simple. Choose the number of random paragraphs you'd like to see and click the button. Your chosen number of paragraphs will instantly appear."

	ch:= make(chan []string)

    go Send(st,ch)
    go receive(st,ch)

	res:= <-ch

	fmt.Println(res)

}