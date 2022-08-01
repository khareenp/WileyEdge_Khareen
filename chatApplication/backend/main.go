//

package main

import (
	"fmt"
	"go-chat/pkg/websocket" //module import
	"io"
	"net/http"
)
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    io.WriteString(w, `{"alive": true}`)
}

func serveWs(lobby *websocket.Lobby, w http.ResponseWriter, r *http.Request) {
    
    fmt.Println("WebSocket Endpoint Hit")
    conn, err := websocket.Upgrade(w, r)
    if err != nil {
        fmt.Fprintf(w, "%+v\n", err)
    }

    client := &websocket.Client{
        Conn: conn,
        Lobby: lobby,
    }

    lobby.Register <- client
    client.Read()
}

func setupRoutes() {
    lobby := websocket.NewLobby()//creates memory location
    go lobby.Start()//1)Server is in listen mode, 2)lobby.Register <- client => 3) client.Read()

    http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
        serveWs(lobby, w, r)//Read from client
        http.HandleFunc("/ws",HealthCheckHandler)
    })
}

func main() {
    fmt.Println("Chat App v1")
    setupRoutes()
    
    http.ListenAndServe(":8080", nil)
}