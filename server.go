package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/websocket"
)

func main() {
    http.HandleFunc("/", wshandler)
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

var (
    wsupgrader = websocket.Upgrader{
        ReadBufferSize:  1024,
        WriteBufferSize: 1024,
        CheckOrigin: func(r *http.Request) bool {
            return true
        },
    }
    nextConnId int = 0
    clients = make(map[int]*websocket.Conn)
)

func wshandler(w http.ResponseWriter, r *http.Request) {
    ws, err := wsupgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println("Failed to set websocket upgrade: %+v", err)
        return
    }

    clients[nextConnId] = ws;
    nextConnId++
    broadcastClientCount()
    for {
        t, msg, err := ws.ReadMessage()
        if err != nil {
            break
        }
        ws.WriteMessage(t, msg)
    }

}

type clientCountMessage struct {
    Action string `json:"action"`
    Data int `json:"data"`
}

func broadcastClientCount() {
    message := clientCountMessage{"clientcount", len(clients)}
    for _, ws := range clients {
        ws.WriteJSON(message)
    }
}





