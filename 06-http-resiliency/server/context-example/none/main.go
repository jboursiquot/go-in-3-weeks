package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func greetHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling greeting request")
	defer log.Println("Handled greeting request")

	completeAfter := time.After(10 * time.Second)

	for {
		select {
		case <-completeAfter:
			fmt.Fprintln(w, "Hello Gopher!")
			return
		default:
			time.Sleep(1 * time.Second)
			log.Println("Greetings are hard. Thinking...")
		}
	}
}

func main() {
	http.HandleFunc("/", greetHandler)
	log.Println("Server listening on :8080...")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
