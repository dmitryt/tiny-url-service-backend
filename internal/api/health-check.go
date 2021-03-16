package api

import (
	"encoding/json"
	"net/http"
)

func (p *API) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	dummyResponse := DummyResponse{true}

	content, err := json.Marshal(dummyResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(content)
}
