package middlewares

import (
	"log/slog"
	"net/http"
	"strings"
	"sync"

	"golang.org/x/time/rate"
)

var (
	clients = make(map[string]*rate.Limiter)
	mu      sync.Mutex
)

func getLimiter(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	limiter, exists := clients[ip]
	if !exists {
		limiter = rate.NewLimiter(3, 5)
		clients[ip] = limiter
	}
	return limiter
}

func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		host := strings.Split(r.RemoteAddr, ":")
		// Get IP from X-Forwarded-For header first, fallback to RemoteAddr
		var ip string
		if forwardedFor := r.Header.Get("X-Forwarded-For"); forwardedFor != "" {
			// X-Forwarded-For can contain multiple IPs, take the first one
			ips := strings.Split(forwardedFor, ",")
			ip = strings.TrimSpace(ips[0])
		} else {
			// Fallback to RemoteAddr and remove port
			if len(host) >= 1 {
				ip = strings.Join(host[0:len(host)-1], ":")
			} else {
				ip = host[0]
			}
		}

		slog.Info("Rate limit middleware", slog.String("ip", ip))
		limiter := getLimiter(ip)

		if !limiter.Allow() {
			slog.Warn("Rate limit exceeded", slog.String("ip", ip))
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
