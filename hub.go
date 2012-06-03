package main

type hub struct {
    listener *connection
    listening bool
    command chan string
    nowplaying chan string
    connect chan *connection
    disconnect chan *connection
}

var h = hub{
    command: make(chan string),
    nowplaying: make(chan string),
    connect: make(chan *connection),
    listening: false,
}

func (h *hub) run() {
    for {
        select {
        case <- h.disconnect:
            if h.listening {
                h.listener.send <- "disconnect plz"
                close(h.listener.send)
                h.listener.ws.Close()
                h.listening = false
            }
        case c := <- h.connect:
            if !h.listening {
                h.listener = c
                h.listening = true
            } else {
                close(c.send)
                c.ws.Close()
            }
        case  m := <- h.command:
            if h.listening {
                h.listener.send <- m 
            }
        }
    }
}
