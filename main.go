package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/api", apiHandler)
	mux.HandleFunc("/api/time", timeHandler)
	mux.HandleFunc("/api/greet", greetHandler)

	wrappedMux := loggingMiddleware(mux)

	log.Println("Starting server on :80")
	err := http.ListenAndServe(":80", wrappedMux)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	respondJSON(w, http.StatusOK, Response{
		Message: "Welcome to Simple Go Backend",
	})
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, Response{
		Message: "API endpoint",
		Data: map[string]interface{}{
			"version": "1.0",
			"routes": []string{
				"/",
				"/api",
				"/api/time",
				"/api/greet?name=YOUR_NAME",
			},
		},
	})
}
func greetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	name := r.URL.Query().Get("name")
	if name == "" {
		respondJSON(w, http.StatusBadRequest, Response{
			Message: "Missing 'name' parameter",
		})
		return
	}

	respondJSON(w, http.StatusOK, Response{
		Message: fmt.Sprintf("Hello, %s!", name),
		Data: map[string]string{
			"name":    name,
			"request": r.URL.RawQuery,
		},
	})
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	respondJSON(w, http.StatusOK, Response{
		Message: "Current server time",
		Data: map[string]interface{}{
			"timestamp": time.Now().Unix(),
			"utc":       time.Now().UTC(),
		},
	})
}

func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}
