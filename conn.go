package main

import (
    "code.google.com/p/go.net/websocket"
)

type connection struct {
    ws *websocket.Conn
    send chan string
}

func (c *connection) reader() {
    for {
        var message string
        err := websocket.Message.Receive(c.ws, &message)
        if err != nil {
            break
        } else if message != "ka" {
            h.command <- "np" + message
        }
    }
    h.disconnect <- c
}

func (c *connection) writer() {
    for message := range c.send {
        err := websocket.Message.Send(c.ws, message)
        if err != nil {
            break
        }
    }
    h.disconnect <- c
    c.ws.Close()
}

func wsHandler(ws *websocket.Conn) {
    c := &connection{send: make(chan string, 256), ws: ws}
    h.connect <- c
    defer func() { h.disconnect <- c }()
    go c.writer()
    c.reader()
}
