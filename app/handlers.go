package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	RespondJSON(w, http.StatusOK, Response{
		Message: "Welcome to Simple Go Backend",
	})
}

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, Response{
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

func GreetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	name := r.URL.Query().Get("name")
	if name == "" {
		RespondJSON(w, http.StatusBadRequest, Response{
			Message: "Missing 'name' parameter",
		})
		return
	}

	RespondJSON(w, http.StatusOK, Response{
		Message: fmt.Sprintf("Hello, %s!", name),
		Data: map[string]string{
			"name":    name,
			"request": r.URL.RawQuery,
		},
	})
}

func TimeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	RespondJSON(w, http.StatusOK, Response{
		Message: "Current server time",
		Data: map[string]interface{}{
			"timestamp": time.Now().Unix(),
			"utc":       time.Now().UTC(),
		},
	})
}

func RespondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
