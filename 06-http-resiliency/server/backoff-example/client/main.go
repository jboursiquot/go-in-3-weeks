package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	var err error
	var requestCount int = 25
	var successCount int
	var throttledCount int

	url := "http://localhost:8080"

	successCount, throttledCount, err = get(url, requestCount)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Printf("Summary (no retry/backoff):  Success: %d, Throttled: %d", successCount, throttledCount)

	// successCount, throttledCount, err = getWithBackoff(url, requestCount)
	// if err != nil {
	// 	log.Println(err)
	// 	os.Exit(1)
	// }
	// log.Printf("Summary (with retry/backoff):  Success: %d, Throttled: %d", successCount, throttledCount)

}

func get(url string, count int) (successCount int, throttledCount int, err error) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < count; i++ {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000))) // Simulate network latency

		var req *http.Request
		req, err = http.NewRequest("GET", url, nil)
		if err != nil {
			return
		}

		req.Header.Set("X-Request-Id", strconv.Itoa(i)) // Add a request id to simulate tracking of requests
		client := &http.Client{}
		var res *http.Response
		res, err = client.Do(req)
		if err != nil {
			return
		}

		res.Body.Close()

		if res.StatusCode == 200 {
			successCount++
		}
		if res.StatusCode == 429 {
			throttledCount++
		}
		log.Printf("%s: %d (Request ID: %v)", url, res.StatusCode, i)
	}
	return
}

func getWithBackoff(url string, count int) (successCount int, throttledCount int, err error) {
	var backoffSchedule = []time.Duration{
		100 * time.Millisecond,
		200 * time.Millisecond,
		400 * time.Millisecond,
		800 * time.Millisecond,
		1000 * time.Millisecond,
	}

	for i := 0; i < count; i++ {
		var res *http.Response

		for _, backoff := range backoffSchedule {
			var req *http.Request
			req, err = http.NewRequest("GET", url, nil)
			if err != nil {
				return
			}

			req.Header.Set("X-Request-Id", strconv.Itoa(i)) // Add a request id to simulate tracking of requests
			client := &http.Client{}
			res, err = client.Do(req)
			if err != nil {
				return
			}

			if res.StatusCode == 200 {
				successCount++
				break
			}
			if res.StatusCode == 429 {
				throttledCount++
			}
			log.Printf("Got status code %d for request %v, back off %v", res.StatusCode, i, backoff)
			time.Sleep(backoff)
		}

		res.Body.Close()
		log.Printf("%s: %d", url, res.StatusCode)
	}
	return
}
