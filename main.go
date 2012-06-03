package main

import (
    "code.google.com/p/go.net/websocket"
    "net/http"
    "fmt"
)

func apiHandler(c http.ResponseWriter, req *http.Request) {
    fmt.Fprint(c, "commando" + req.URL.Path[:1])
}


func main() {
    go h.run()
    http.HandleFunc("/", apiHandler)
    http.Handle("/ws", websocket.Handler(wsHandler))
    err := http.ListenAndServe(":9002", nil)
    if err != nil {
        fmt.Println("het is niet gelukt.. helaas")
    }
}
