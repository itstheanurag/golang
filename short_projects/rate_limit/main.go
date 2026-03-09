package main

import (
	"fmt"
	"net/http"
	"rate-limit/middleware"
	"rate-limit/ratelimit"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	fmt.Println("RATE LIMITING IN GOLANG")

	limiter := ratelimit.NewFixedWindowRateLimiter(10, 20*time.Second)

	http.Handle("/", middleware.RateLimiterMiddleware(limiter, http.HandlerFunc(Handler)))

	fmt.Printf("Starting go server on the Port: %d\n", 8080)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server failed:", err)
	}
}
