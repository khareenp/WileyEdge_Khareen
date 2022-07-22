package clients

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

var port = flag.String("port", "8000", "the port number")
var tz = flag.String("tz", "8000", "time zone")

// call like ./clock1 -port 8010 -tz=Asia/Tokyo
func clock1() {
	flag.Parse()
	os.Setenv("TZ", *tz)
	listener, err := net.Listen("tcp", "localhost:"+*port)
	fmt.Printf("Listening on port %s\n", *port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle connections concurrently
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		fmt.Println("Received connection from: ", c.RemoteAddr())
		time.Sleep(1 * time.Second)
	}
}