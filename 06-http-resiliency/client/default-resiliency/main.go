package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Making request...")
	res, err := http.Get("http://localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res.Status)
	defer res.Body.Close()
}
