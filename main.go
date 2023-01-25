package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const (
	addr       = ":8080"
	appVersion = "v1.0.0"
)

func version(w http.ResponseWriter, req *http.Request) {
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "App version: %s\n", appVersion)
	fmt.Fprintf(w, "Hostname: %s\n", hostname)
	fmt.Fprintf(w, "Time: %s\n", time.Now().Local())
}

func main() {
	fmt.Printf("Server started at %s\n", addr)
	http.HandleFunc("/", version)
	http.ListenAndServe(addr, nil)
}
