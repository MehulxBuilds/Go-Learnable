package external

import (
	"go-http/internal/shared"
	"net/http"
	"time"
)

func ExternalHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		shared.WriteJSON(w, http.StatusMethodNotAllowed, map[string]any{
			"ok":    false,
			"error": "Only GET method is allowed",
		})
		return
	}

	data, err := shared.ApiCallerWithUnmarshal()
	if err != nil {
		shared.WriteJSON(w, http.StatusBadGateway, map[string]any{
			"ok":    false,
			"error": "Failed to fetch data",
		})
		return
	}

	shared.WriteJSON(w, http.StatusOK, map[string]any{
		"ok":        true,
		"timeStamp": time.Now().UTC(),
		"external": map[string]any{
			"source": "Catfact.ninja",
			"fact":   data.Fact,
			"length": data.Length,
		},
	})
}