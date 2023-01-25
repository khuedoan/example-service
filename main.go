package main

import (
    "fmt"
    "net/http"
)

const (
    addr = ":8080"
    appVersion = "v1.0.0"
)

func version(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "App version: %s\n", appVersion)
}

func main() {
    fmt.Printf("Server started at %s\n", addr)
    http.HandleFunc("/", version)
    http.ListenAndServe(addr, nil)
}
