package main

import (
    "net/http"
    "fmt"
)


func main() {

    //http.HandleFunc("/protobuf", ProtobufHandler)
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello world")
    })
    http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Only a test")
    })

    fmt.Println("Connected to mongodb, wait for connections")
    http.ListenAndServe(":8080", nil)
}
