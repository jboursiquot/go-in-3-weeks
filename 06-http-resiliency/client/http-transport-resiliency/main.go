package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	tr := &http.Transport{
		TLSHandshakeTimeout: 10 * time.Second,
		MaxIdleConns:        10,
		//...
	}

	client := &http.Client{
		Timeout:   time.Second * 3, // Timeout for the entire request
		Transport: tr,
	}

	log.Println("Making request...")
	res, err := client.Get("http://localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	log.Println(res.Status)
}
