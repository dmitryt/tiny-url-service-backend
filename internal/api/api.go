package api

import (
	"github.com/gorilla/mux"
)

type API struct {
}

type DummyResponse struct {
	OK bool
}

func New() *API {
	return &API{}
}

func (p *API) Handle(r *mux.Router) {
	p.HandleLinks(r.PathPrefix("/links").Subrouter());
	p.HandleHealthCheck(r.PathPrefix("/health-check").Subrouter());
}
