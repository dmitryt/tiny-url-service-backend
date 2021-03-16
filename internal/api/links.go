package api

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrImageFetch         = errors.New("error during fetching the image")
	ErrImageResize        = errors.New("error during resizing the image")
	ErrInvalidURI         = errors.New("invalid URI. Expected format is: /<method>/<width>/<height>/<external url>")
	ErrImageCopyFromCache = errors.New("error during copying the image from cache")
)

func (p *API) LinksHandler(w http.ResponseWriter, r *http.Request) {
	dummyResponse := DummyResponse{true}

	content, err := json.Marshal(dummyResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(content)
}
