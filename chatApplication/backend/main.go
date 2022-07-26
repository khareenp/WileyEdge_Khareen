package main

import (
	"fmt"
	"go-chat/pkg/websocket" //module import
	"net/http"
)

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
    lobby := websocket.NewLobby()
    go lobby.Start()

    http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
        serveWs(lobby, w, r)
    })
}

func main() {
    fmt.Println("Distributed Chat App v0.01")
    setupRoutes()
    http.ListenAndServe(":8080", nil)
}