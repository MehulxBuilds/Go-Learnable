package main

import (
	"encoding/json"
	"fmt"
	"go-http/internal/external"
	"go-http/internal/greet"
	"go-http/internal/shared"
	"net/http"
	"strings"
	"time"
)

type TestRequest struct {
	Name string `json:"name"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only Get is allowed", http.StatusMethodNotAllowed)
		return
	}
	_,_ = w.Write([]byte("Welcome to the Go HTTP Server!"))
}

func successHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	res := map[string]any{
		"ok":       true,
		"message":  "JSON encode successfull",
		"datetime": time.Now().UTC(),
	}

	_ = json.NewEncoder(w).Encode(res)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		shared.WriteJSON(w, http.StatusMethodNotAllowed, map[string]any{
			"ok":    false,
			"error": "Only post is allowed",
		})
		return
	}

	defer r.Body.Close()

	var req TestRequest

	dec := json.NewDecoder(r.Body)

	if err := dec.Decode(&req); err != nil {
		shared.WriteJSON(w, http.StatusBadRequest, map[string]any{
			"ok":    false,
			"error": "Invalid json format",
		})
		return
	}

	req.Name = strings.TrimSpace(req.Name)

	if req.Name == "" {
		shared.WriteJSON(w, http.StatusBadRequest, map[string]any{
			"ok":    false,
			"error": "name must not be empty!",
		})
		return
	}

	shared.WriteJSON(w, http.StatusOK, map[string]any{
		"ok":        true,
		"data":      req.Name,
		"timeStamp": time.Now().UTC(),
	})
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/ok", successHandler)
	http.HandleFunc("/test", testHandler)
	http.HandleFunc("/greet", greet.HandleGreet)
	http.HandleFunc("/external", external.ExternalHandler)	

	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}