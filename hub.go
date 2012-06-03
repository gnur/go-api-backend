package main

type hub struct {
    listener *conection
    listening bool
    command chan string
    nowplaying chan string
    starting chan bool
}

var h = hub{
    command: make(chan string),
    nowplaying: make(chan string),
    listening: false
}

func (h *hub) run() {
    for {
        if b := <-h.starting {
            listening = b
        }
        if  m := <-h.command && h.listening {
            c.send <- m 
        }
    }
}
