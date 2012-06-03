package main

import (
    "code.google.com/p/go.net/websocket"
    "net/http"
)

func main() {
    go h.run()
    http.HandleFunc("/", apiHandler)
    http.Handle("/ws", websocket.Handler(wsHandler))
    err := http.ListenAndServe(":9002", nil)
    if err != nil {
        fmt.Println("het is niet gelukt.. helaas")
    }
}
