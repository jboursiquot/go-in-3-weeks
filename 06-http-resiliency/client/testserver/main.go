package main

import (
	"fmt"
	"net/http"
	"time"
)

type greeter struct{}

func (h *greeter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second * 60)
	fmt.Fprintf(w, "Greetings!")
}

func main() {
	fmt.Println("Listening on :8080...")
	http.ListenAndServe("127.0.0.1:8080", &greeter{})
}
