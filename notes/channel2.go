package notes

import (
	"fmt"
	//"sync"
	//"time"
)
//substitute for main function
func channel2() {
	send:=make(chan int ,1)
    reciev:=make(chan int ,1)
    go sender(send ,reciev)
    go reciever(reciev ,send )
    send<-1
    select{ }//allow both routines ot execute full cycle. if removed will have to use waitgroup
} // channel2 () over 
func sender( snd <-chan int , recv chan <- int   ){ // (snd <-chan int(Reading ) , recv chan <- int (writing ) )
    for {
                inform:=<-snd  // first time inform is 1 
                fmt.Println("information from  Send " , inform)
                //time.Sleep(1*time.Second)
                recv <-inform+1
    }
}
func reciever( recv <- chan int , snd chan<- int ){

            for {
                inform:= <-recv // inform is recieved 2
                 fmt.Println("information from  Reciever " , inform)
                //time.Sleep(1*time.Second)
                snd <- inform+1
                }
}