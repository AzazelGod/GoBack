package tests

import (
	"encoding/json"
	"goback/app"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.HomeHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response app.Response
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}

	expected := app.Response{
		Message: "Welcome to Simple Go Backend",
	}

	if response.Message != expected.Message {
		t.Errorf("handler returned unexpected message: got %v want %v",
			response.Message, expected.Message)
	}
}

func TestApiHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.ApiHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response app.Response
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}

	expected := app.Response{
		Message: "API endpoint",
		Data: map[string]interface{}{
			"version": "1.0",
			"routes": []interface{}{
				"/",
				"/api",
				"/api/time",
				"/api/greet?name=YOUR_NAME",
			},
		},
	}

	if response.Message != expected.Message {
		t.Errorf("handler returned unexpected message: got %v want %v",
			response.Message, expected.Message)
	}

	// Проверка данных
	responseData, ok := response.Data.(map[string]interface{})
	if !ok {
		t.Fatal("response data is not a map")
	}

	expectedData := expected.Data.(map[string]interface{})
	if responseData["version"] != expectedData["version"] {
		t.Errorf("handler returned unexpected version: got %v want %v",
			responseData["version"], expectedData["version"])
	}

	// Проверка маршрутов
	responseRoutes, ok := responseData["routes"].([]interface{})
	if !ok {
		t.Fatal("routes data is not a slice")
	}

	expectedRoutes := expectedData["routes"].([]interface{})
	if len(responseRoutes) != len(expectedRoutes) {
		t.Errorf("handler returned unexpected number of routes: got %v want %v",
			len(responseRoutes), len(expectedRoutes))
	}

	for i, route := range responseRoutes {
		if route != expectedRoutes[i] {
			t.Errorf("handler returned unexpected route: got %v want %v",
				route, expectedRoutes[i])
		}
	}
}

func TestGreetHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/greet?name=John", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.GreetHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response app.Response
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}

	if response.Message != "Hello, John!" {
		t.Errorf("handler returned unexpected message: got %v want %v",
			response.Message, "Hello, John!")
	}
}

func TestTimeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/time", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.TimeHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response app.Response
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}

	if response.Message != "Current server time" {
		t.Errorf("handler returned unexpected message: got %v want %v",
			response.Message, "Current server time")
	}
}
