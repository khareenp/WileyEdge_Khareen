package websocket

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
    ID   string
    Conn *websocket.Conn //pointer to a websocket.Conn object
    Lobby *Lobby //a pointer to the Lobby which this client will be part of
}

type Message struct {
    Type int    `json:"type"`
    Body string `json:"body"`
}

// Read() method which will constantly listen in for new messages 
// coming through on this Client’s websocket connection.
func (c *Client) Read() {
    defer func() {
        c.Lobby.Unregister <- c
        c.Conn.Close()
    }()

    for {
        messageType, p, err := c.Conn.ReadMessage()
        if err != nil {
            log.Println(err)
            return
        }
        message := Message{Type: messageType, Body: string(p)}
        c.Lobby.Broadcast <- message// message is sent to all clients
        fmt.Printf("Message Received: %+v\n", message)
    }
}
// If there are any messages, it will pass these messages to the Pool’s Broadcast channel 
// which subsequently broadcasts the received message to every client within the pool.

