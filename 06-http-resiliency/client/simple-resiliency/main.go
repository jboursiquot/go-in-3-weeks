package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: time.Second * 3,
	}

	log.Println("Making request...")

	res, err := client.Get("http://localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	log.Println(res.Status)
}
