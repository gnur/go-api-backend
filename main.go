package main

import (
    "code.google.com/p/go.net/websocket"
    "net/http"
    "fmt"
    "strings"
)

func apiHandler(c http.ResponseWriter, req *http.Request) {
    var commando, parameter string
    c.Header().Set("access-control-allow-origin", "http://gnur.nl:8080")
    commands := strings.Split(req.URL.Path[1:], "/")
    if len(commands) > 1 {
        commando = commands[1]
    } else {
        commando = "nada"
    }
    if len(commands) > 2 {
        parameter = commands[2]
    } else {
        parameter = "nada"
    }
    if parameter == "nada" {
        switch commando {
        case "playpause":
            h.command <- "playPause();"
        case "next":
            h.command <- "nextSong();"
        case "play":
            h.command <- "resumeSong();"
        case "pause":
            h.command <- "resumeSong();"
        case "clearfilter":
            h.command <- "filterSongs(\"\");"
        }
    } else {
        switch commando {
        case "playsong":
            h.command <- "playSong('" + parameter + "');"
        case "quesong":
            h.command <- "queSong('" + parameter + "');"
        case "filter":
            h.command <- "filterSongs('" + parameter + "');"
        case "volume":
            h.command <- "setVolume('" + parameter + "');"
        }
    }
    fmt.Println(commando + ": " + parameter)

    fmt.Fprint(c, "commando " + commando + "\nparameter " + parameter)
    
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
