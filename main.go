package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	addr := ":8080"
	http.HandleFunc("/", MyHandler)
	log.Println("Server started on port", addr)
	log.Fatal(http.ListenAndServe(addr, nil))

}

func MyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}
