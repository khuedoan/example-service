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

func version(w http.ResponseWriter, req *http.Request) {
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

	http.HandleFunc("/", version)
	http.Serve(lis, nil)
}
