package main

import (
	"fmt"
	"net/http"
)

type greeter struct{}

func (h *greeter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Greetings!")
}

func main() {
	http.ListenAndServe("127.0.0.1:8080", &greeter{})
}
