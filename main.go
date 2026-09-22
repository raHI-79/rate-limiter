package main
import (
	"fmt"
	"net"
	"net/http"
	"os"
	"sync"
	"time"
)
type Client struct {
	count int
	time  time.Time
}
var clients = make(map[string]Client)
var mu sync.Mutex
const limit = 5
func getIP(r *http.Request) string {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	return ip
}
func rateLimit(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip := getIP(r)
		mu.Lock()
		client, ok := clients[ip]
		if !ok || time.Since(client.time) >= time.Minute {
			clients[ip] = Client{count: 1, time: time.Now()}
			mu.Unlock()
			next(w, r)
			return
		}
		if client.count >= limit {
			mu.Unlock()
			w.WriteHeader(http.StatusTooManyRequests)
			fmt.Fprint(w, "Too Many Requests! Try again later.")
			return
		}
		client.count++
		clients[ip] = client
		mu.Unlock()
		next(w, r)
	}
}
func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
func request(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Request accepted!")
}
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", home)
	http.HandleFunc("/request", rateLimit(request))
    fmt.Println("Server running on http://localhost:" + port + "/")
	http.ListenAndServe(":"+port, nil)
}