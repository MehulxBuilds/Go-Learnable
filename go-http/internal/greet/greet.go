package greet

import (
	"fmt"
	"net/http"
)

func HandleGreet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only Get is allowed", http.StatusMethodNotAllowed)
		return
	}

	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	greeting := fmt.Sprintf("Hello, %s!", name)
	_,_ = w.Write([]byte(greeting))
}

