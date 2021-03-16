package api

type API struct {
}

type DummyResponse struct {
	OK2212211 bool
}

func New() *API {
	return &API{}
}
