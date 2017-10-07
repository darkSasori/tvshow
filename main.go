package main

import (
    "net/http"
    "fmt"
    "log"
	"os"
)

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Version 1.1.3")
    fmt.Fprintf(w, os.Getenv("HOSTNAME"))
}

func main() {
    http.HandleFunc("/import", ImportHandler)
    http.HandleFunc("/", HelloWorldHandler)

    log.Println("Connected to mongodb, wait for connections on :8080")
    http.ListenAndServe(":8080", nil)
}
