package main

import (
	"fmt"
	"log"
	"net/http"
)

type morningHandler struct{}

func (mh *morningHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Good Morning!")
}

type eveningHandler struct{}

func (mh *eveningHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/evening/tea":
		w.WriteHeader(http.StatusTeapot)
		fmt.Fprintf(w, "Good Evening! Have some tea.")
		return
	case "/evening":
		fmt.Fprintf(w, "Good Evening!")
		return
	}
}

func main() {
	eh := &eveningHandler{}
	mh := &morningHandler{}

	mux := http.NewServeMux()
	mux.Handle("/evening/tea", eh)
	mux.Handle("/evening", eh)
	mux.Handle("/morning", mh)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Greetings!")
	})

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", mux))
}
