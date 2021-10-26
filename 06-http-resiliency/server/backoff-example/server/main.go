package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/time/rate"
)

var limiter *rate.Limiter

// limit is middleware that rate limits requests
func limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			log.Printf("Rate limit exceeded (Request ID: %v)", r.Header.Get("X-Request-Id"))
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

const (
	defaultPort  = 8080
	defaultRate  = 1
	defaultBurst = 3
)

func main() {
	port := fmt.Sprintf(":%d", *flag.Int("port", defaultPort, "port (int)"))
	r := flag.Float64("rate", defaultRate, "rate limit (float)")
	b := flag.Int("burst", defaultBurst, "burst limit (int)")
	flag.Parse()
	limiter = rate.NewLimiter(rate.Limit(*r), *b)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(http.StatusText(http.StatusOK)))
	})

	log.Printf("Server ready on %s with allowed rate of %v req/s and burst of %v reqs...", port, *r, *b)
	http.ListenAndServe(port, limit(mux))
}
