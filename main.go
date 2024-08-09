package main

import (
    "flag"
    "fmt"
    "log"
    "net/http"
)

func main() {
    var messageFlag = flag.String("message", "Hello World!", "Message of mock server to be display")
    var portFlag = flag.Uint("port", 80, "Port to run server on")

    flag.Parse()

    var message = []byte(*messageFlag)
    var port = fmt.Sprint(":", *portFlag)

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        log.Print("Servered\n\n", r, "\n\n")
        var _, err = w.Write(message)
        if err != nil {
            log.Println(err)
        }
    })

    log.Println("Server is starting on", port, "to tell the world:\n", *messageFlag)
    log.Fatalln(http.ListenAndServe(port, nil))
}
