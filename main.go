package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ServiceWeaver/weaver"
)

const (
	addr       = ":8080"
	appVersion = "v1.0.0"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!\n", r.URL.Query().Get("name"))
}

func info(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "App version: %s\n", appVersion)
	fmt.Fprintf(w, "Hostname: %s\n", hostname)
	fmt.Fprintf(w, "Time: %s\n", time.Now().Local())
}

func main() {
	root := weaver.Init(context.Background())

	opts := weaver.ListenerOptions{
		LocalAddress: addr,
	}

	lis, err := root.Listener("example", opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("example listener available on %v\n", lis)

	http.HandleFunc("/", info)
	http.HandleFunc("/info", info)
	http.HandleFunc("/hello", hello)
	http.Serve(lis, nil)
}
