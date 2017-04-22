package main

import (
    "net/http"
    "fmt"
)

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello world")
}

func main() {

    http.HandleFunc("/import", ImportHandler)

    http.HandleFunc("/", HelloWorldHandler)

    fmt.Println("Connected to mongodb, wait for connections")
    http.ListenAndServe(":8080", nil)
}
