package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleRequest() {
	http.HandleFunc("/", rootFunc)
	log.Fatal(http.ListenAndServe(":8088", nil))
}

func rootFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from root")
	fmt.Println("In rootFunc")
}

func main() {
	handleRequest()
}
