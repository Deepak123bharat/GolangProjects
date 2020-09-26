// youtube-stats/main.go
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/deepak123bharat/youtube-stats/websocket"
)

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World")
}

func stats(w http.ResponseWriter, r *http.Request) {
    ws, err := websocket.Upgrade(w, r)
    if err != nil {
        fmt.Fprintf(w, "%+v\n", err)
    }
    go websocket.Writer(ws)
}

func setupRoutes() {
    port := os.Getenv("PORT")
    http.Handle("/", http.FileServer(http.Dir("./")))
    http.HandleFunc("/stats", stats)
    log.Println("Listening on :"+ string(port))
    log.Fatal(http.ListenAndServe(":"+"8080", nil))
}

func main() {
    
    fmt.Println("YouTube Subscriber Monitor")
    setupRoutes()
}