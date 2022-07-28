package websocket

import "fmt"

// struct which will contain all of the channels we need for concurrent communication,
// as well as a map of clients.
type Lobby struct {
    Register   chan *Client
    Unregister chan *Client
    Clients    map[*Client]bool
    Broadcast  chan Message
}

func NewLobby() *Lobby {//firt function to be called
    return &Lobby{
        Register:   make(chan *Client),
        Unregister: make(chan *Client),
        Clients:    make(map[*Client]bool),
        Broadcast:  make(chan Message),
    }
}

func (lobby *Lobby) Start() {
    for {
        select {
			// Register - Our register channel will send out New User Joined... 
			// to all of the clients within this lobby when a new client connects.
        case client := <-lobby.Register:
            lobby.Clients[client] = true
            fmt.Println("Size of Connection Lobby: ", len(lobby.Clients))
            for client, _ := range lobby.Clients {
                fmt.Println(client)
                client.Conn.WriteJSON(Message{Type: 1, Body: "New User Joined..."})            
            }
            break

		// Unregister - Will unregister a user and notify the lobby when a client disconnects.
        case client := <-lobby.Unregister:
            delete(lobby.Clients, client)
            fmt.Println("Size of Connection Lobby: ", len(lobby.Clients))
            for client, _ := range lobby.Clients {
                client.Conn.WriteJSON(Message{Type: 1, Body: "User Disconnected..."})
            }
            break
		

		// Broadcast - a channel which, when it is passed a message, will loop through all clients 
		// in the lobby and send the message through the socket connection.
        case message := <-lobby.Broadcast:
            fmt.Println("Sending message to all clients in Lobby")
            for client, _ := range lobby.Clients {
                if err := client.Conn.WriteJSON(message); err != nil {
                    fmt.Println(err)
                    return
                }
            }
        }
    }
}