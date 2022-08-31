package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strings"
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ts.URL, nil)
	if err != nil {
		log.Fatalln(err)
	}

	t := time.Now()
	log.Println("Sending request...")

	log.Printf("Client >> Making request to test server: %s", ts.URL)
	r, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Client >> Error: %s", err)
		log.Println("Client >> Moving on...")
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	msg := strings.TrimSpace(string(b))
	log.Printf("Client >> Received response %q in %v", msg, time.Since(t))
}
