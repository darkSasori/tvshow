package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/user"
)

func middlewareLogger(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s from %s\n", r.Method, r.URL.Path, r.RemoteAddr)
		fn(w, r)
	}
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Version 1.1.3\n")
	fmt.Fprintf(w, os.Getenv("HOSTNAME"))
}

func main() {
	http.HandleFunc("/import", middlewareLogger(ImportHandler))
	http.HandleFunc("/", middlewareLogger(HelloWorldHandler))

	if user, err := user.Current(); err == nil {
		log.Printf("Using '%s'\n", user.Name)
	}
	log.Printf("MongoDB: '%s'\n", TVSHOW_MONGODB)
	log.Println("Wait for connections on :8080")
	http.ListenAndServe(":8080", nil)
}
