package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleRequest() {
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from root")
	log.Println("In Root Handler")
}

func main() {
	handleRequest()
}
