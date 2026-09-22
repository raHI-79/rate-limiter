package main
import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
)
type Client struct {
	Count int
	Time  time.Time
}
var clients = make(map[string]Client)
var mu sync.Mutex
const limit = 5
func getIP(r *http.Request) string {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return strings.TrimSpace(r.RemoteAddr)
	}
	return ip
}
func rateLimit(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip := getIP(r)
		mu.Lock()
		client, exists := clients[ip]
		if !exists || time.Since(client.Time) >= time.Minute {
			clients[ip] = Client{
				Count: 1,
				Time:  time.Now(),
			}
			mu.Unlock()
			next(w, r)
			return
		}
		if client.Count >= limit {
			mu.Unlock()
			w.WriteHeader(http.StatusTooManyRequests)
			fmt.Fprintln(w, "Too Many Requests! Try again later.")
			return
		}
		client.Count++
		clients[ip] = client
		mu.Unlock()
		next(w, r)
	}
}
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, "Request accepted!")
}
func main() {
	http.HandleFunc("/", rateLimit(home))
	fmt.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}