package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Link struct {
	ID    string `json:"_id"`
	Alias string `json:"alias"`
	URL   string `json:"url"`
}

type LinkPayload struct {
	Value string `json:"value"`
}

var (
	ErrFileCreate       = errors.New("error during creating the file")
	ErrFileRead         = errors.New("error during reading the file")
	ErrFileWrite        = errors.New("error during writing the file")
	ErrContentUnmarshal = errors.New("error during unmarshaling the content")
	ErrContentMarshal   = errors.New("error during marshaling the content")
)

var storagePath = filepath.Join("fixtures", "links.json")

func logErrorIfExists(ctxErr, err error) {
	if err != nil {
		log.Error().Msgf("%s: %s", ctxErr, err)
	}
}

func readContentOrCreateFile() []byte {
	content := []byte("[]")
	if _, err := os.Stat(storagePath); os.IsNotExist(err) {
		_, err := os.Create(storagePath)
		logErrorIfExists(ErrFileCreate, err)
	}
	content, err := ioutil.ReadFile(storagePath)
	logErrorIfExists(ErrFileRead, err)
	return content
}

func readLinksFromFile() (result []Link) {
	content := readContentOrCreateFile()
	err := json.Unmarshal(content, &result)
	logErrorIfExists(ErrContentUnmarshal, err)

	return
}

func writeLinksToFile(data []Link) {
	content, err := json.Marshal(data)
	logErrorIfExists(ErrContentMarshal, err)
	if err == nil {
		err = ioutil.WriteFile(storagePath, content, 0644)
		logErrorIfExists(ErrFileWrite, err)
	}
	return
}

func getLinksHandler(w http.ResponseWriter, r *http.Request) {
	content := readContentOrCreateFile()
	if len(content) == 0 {
		content = []byte("[]")
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(content)
}

func addLinkHandler(w http.ResponseWriter, r *http.Request) {
	var data LinkPayload
	var newLink Link
	result := []byte("{}")
	_ = json.NewDecoder(r.Body).Decode(&data)
	if data.Value != "" {
		links := readLinksFromFile()
		newLink = Link{ID: primitive.NewObjectID().Hex(), Alias: fmt.Sprintf("/%s", uuid.New().String()), URL: data.Value}
		writeLinksToFile(append(links, newLink))
		result, _ = json.Marshal(newLink)
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(result)
}

func deleteLinkHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	linkId := -1
	if _, ok := vars["id"]; ok {
		links := readLinksFromFile()
		for idx, link := range links {
			if link.ID == vars["id"] {
				linkId = idx
				break
			}
		}
		if linkId != -1 {
			writeLinksToFile(append(links[:linkId], links[linkId+1:]...))
		}
	}
	if linkId == -1 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("{\"OK\":true}"))
	}
}

func (p *API) HandleLinks(r *mux.Router) {
	r.HandleFunc("", getLinksHandler).Methods("GET")
	r.HandleFunc("", addLinkHandler).Methods("POST")
	r.HandleFunc("/{id}", deleteLinkHandler).Methods("DELETE")
}
