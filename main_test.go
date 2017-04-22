package main

import (
    "net/http"
    "net/http/httptest"
    "io/ioutil"
    "testing"
)

func TestHelloWorldHandler(t *testing.T) {
    server := httptest.NewServer(http.HandlerFunc(HelloWorldHandler))
    defer server.Close()

    res, err := http.Get(server.URL)
    if err != nil {
        t.Fatal(err)
    }

    if res.StatusCode != 200 {
        t.Fatalf("Received non-200 response: %d\n", res.StatusCode)
    }

    data, err := ioutil.ReadAll(res.Body)
    if err != nil {
        t.Fatal(err)
    }

    if string(data) != "Hello world" {
        t.Errorf("Expected 'Hello world', received '%s'", data)
    }
}
