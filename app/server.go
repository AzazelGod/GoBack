package app

import (
	"log"
	"net/http"
	"time"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/api", ApiHandler)
	mux.HandleFunc("/api/time", TimeHandler)
	mux.HandleFunc("/api/greet", GreetHandler)
	return mux
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}

func StartServer() {
	mux := NewRouter()
	wrappedMux := LoggingMiddleware(mux)

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", wrappedMux)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
