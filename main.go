package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	defaultPort = "8080"
	appVersion  = "v1.0.0"
)

func hello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello, %s!\n", name)
}

func info(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		http.Error(w, "Failed to get hostname", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "App version: %s\n", appVersion)
	fmt.Fprintf(w, "Hostname: %s\n", hostname)
	fmt.Fprintf(w, "Time: %s\n", time.Now().Local())
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.HandleFunc("/", info)
	http.HandleFunc("/info", info)
	http.HandleFunc("/hello", hello)

	addr := fmt.Sprintf(":%s", port)
	fmt.Printf("Server running on %s\n", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
