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
		// Remove port from IP address
		var ip string
		if len(host) >= 1 {
			ip = strings.Join(host[0:len(host)-1], ":")
		} else {
			ip = host[0]
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
