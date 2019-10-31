package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", httpHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func httpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello go http")
}
