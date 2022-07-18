package notes

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func Server() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		handleConn(conn) // handle one connection at a time
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		fmt.Println(c.RemoteAddr())
		time.Sleep(1 * time.Second)
	}
}