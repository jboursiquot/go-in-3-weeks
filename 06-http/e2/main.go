package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Greetings!")
	})

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", h))
}
