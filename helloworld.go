package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	addr := ":8080"
	http.HandleFunc("/", MyHandler)
	log.Println("Server started on port", addr)
	log.Fatal(http.ListenAndServe(addr, nil))

}

type HelloWorldMessage struct {
	Message string
}

func MyHandler(w http.ResponseWriter, r *http.Request) {
	var helloMessage = HelloWorldMessage{"Hello World!"}
	json.NewEncoder(w).Encode(helloMessage)
}
