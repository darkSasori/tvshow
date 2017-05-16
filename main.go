package main

import (
    "net/http"
    "fmt"
)

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Version 1.1.2")
}

func main() {
    http.HandleFunc("/import", ImportHandler)
    http.HandleFunc("/", HelloWorldHandler)

    fmt.Println("Connected to mongodb, wait for connections on :8080")
    http.ListenAndServe(":8080", nil)
}
