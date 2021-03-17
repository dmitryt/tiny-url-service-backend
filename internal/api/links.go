package api

import (
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

type Link struct {
	ID    string `json:"_id"`
	Value string `json:"value"`
}

func getLinksHandler(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile("fixtures/links.json")
	if err != nil {
		log.Fatal().Err(err)
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(content)
}

func (p *API) HandleLinks(r *mux.Router) {
	r.HandleFunc("", getLinksHandler).Methods("GET")
}
