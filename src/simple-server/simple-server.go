package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	log.Printf("Received request from %s\n", req.UserAgent())
	fmt.Fprintf(w, "Server: simple-server\n")
}

func main() {
	http.HandleFunc("/", hello)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}