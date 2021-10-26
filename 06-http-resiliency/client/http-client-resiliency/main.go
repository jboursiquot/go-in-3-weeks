package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"time"
)

func main() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Server >> Request received...[%s] %s", r.Method, r.RequestURI)
		rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
		time.Sleep(time.Duration(rnd.Intn(60)) * time.Second) // Simulate a slow server
		msg := "Hello Gopher!"
		log.Printf("Server >> Sending %q", msg)
		fmt.Fprintln(w, msg)
	}))
	defer ts.Close()

	client := &http.Client{
		Timeout: time.Second * 3,
	}

	log.Println("Making request...")

	res, err := client.Get(ts.URL)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	log.Println(res.Status)
}
