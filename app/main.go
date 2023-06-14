package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloWorldHandler)

	err := http.ListenAndServe(":9000", nil)

	if err != nil {
		log.Printf("Error starting server %v", err)
	}
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Hello Nelson!")

	if err != nil {
		return
	}
}
