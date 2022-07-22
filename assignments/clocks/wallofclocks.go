package clocks

// wall of clocks
import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sync"
)

// ./wallofclocks NewYork=localhost:8010 Tokyo=localhost:8020 London=localhost:8030

var wg sync.WaitGroup

func main() {
	args := os.Args[1:]
	for _,val := range args{
		// split to get the localhost
		// then make connection to that host port
		fmt.Println(val)
		wg.Add(1)
		str:= findLocalhost(val)
		go timeAt(str)

	}
	wg.Wait()
	
}

func findLocalhost(str string) string{
	var returnString string
	for i:=0; i < len(str); i++{
		if string(str[i]) == "="{
			returnString = string(str[i+1:])
			break
		}
	}
	return returnString
}

func shouldCopy2(dst io.Writer, src io.Reader){
	if _, err := io.Copy(dst,src); err != nil{
		log.Fatal(err)
		wg.Done()
	}
}

func timeAt(host string){
	conn, err := net.Dial("tcp",host)
	if err != nil{
		log.Fatal(err)
	}
	shouldCopy2(os.Stdout, conn)
}