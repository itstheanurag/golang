package middleware

import (
	"net/http"
	"strings"
)

type RateLimiter interface {
	Allow(key string) bool
}

func RateLimiterMiddleware(limiter RateLimiter, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		key := extractClientId(r)

		if !limiter.Allow(key) {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func extractClientId(r *http.Request) string {
	if apiKey := r.Header.Get("X-API-Key"); apiKey != "" {
		return "api_key:" + apiKey
	}

	ip := r.RemoteAddr

	if forwardedFor := r.Header.Get("X-Forwarded-For"); forwardedFor != "" {
		ip = strings.Split(forwardedFor, ",")[0]
		ip = strings.TrimSpace(ip)
	}

	if colonIdx := strings.LastIndex(ip, ":"); colonIdx != -1 {
		if strings.Count(ip, ":") == 1 {
			ip = ip[:colonIdx]
		}

	}

	return "client_ip:" + ip
}
