package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	url := "http://localhost:8080"
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req = req.WithContext(ctx)

	t := time.Now()
	log.Println("Sending request...")

	log.Printf("Client >> Making request to test server: %s", url)
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
