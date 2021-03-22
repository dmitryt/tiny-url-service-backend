package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"

	models "github.com/dmitryt/tiny-url-service-backend/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LinkPayload struct {
	Value string `json:"value"`
}

var (
	ErrFileCreate       = errors.New("error during creating the file")
	ErrFileRead         = errors.New("error during reading the file")
	ErrFileWrite        = errors.New("error during writing the file")
	ErrContentUnmarshal = errors.New("error during unmarshaling the content")
	ErrContentMarshal   = errors.New("error during marshaling the content")
	ErrParsingBody      = errors.New("error during parsing the request body")
	ErrValidationAddURL = errors.New("error, incoming data is empty")
)

var storagePath = filepath.Join("fixtures", "links.json")

func logErrorIfExists(ctxErr, err error) error {
	if err != nil {
		log.Error().Msgf("%s: %s", ctxErr, err)

		return ctxErr
	}

	return nil
}

func readLinksFromFile() (result []models.Link, err error) {
	content, err := ioutil.ReadFile(storagePath)
	if err != nil {
		return
	}
	err = json.Unmarshal(content, &result)

	return result, logErrorIfExists(ErrContentUnmarshal, err)
}

func writeLinksToFile(data []models.Link) (err error) {
	content, err := json.Marshal(data)
	if err != nil {
		return logErrorIfExists(ErrContentMarshal, err)
	}
	err = ioutil.WriteFile(storagePath, content, 0o600)

	return logErrorIfExists(ErrFileWrite, err)
}

func getLinksHandler(c *fiber.Ctx) error {
	links, err := readLinksFromFile()
	if err != nil {
		return err
	}

	return c.JSON(links)
}

func addLinkHandler(c *fiber.Ctx) error {
	payload := new(LinkPayload)
	if err := c.BodyParser(payload); err != nil {
		return logErrorIfExists(ErrParsingBody, err)
	}
	if payload.Value == "" {
		return logErrorIfExists(ErrValidationAddURL, nil)
	}
	links, err := readLinksFromFile()
	if err != nil {
		return err
	}
	newLink := models.Link{ID: primitive.NewObjectID().Hex(), Alias: fmt.Sprintf("/%s", uuid.New().String()), URL: payload.Value}
	err = writeLinksToFile(append(links, newLink))
	if err != nil {
		return err
	}

	return c.JSON(newLink)
}

func deleteLinkHandler(c *fiber.Ctx) error {
	linkID := c.Params("id")
	linkIndex := -1
	links, err := readLinksFromFile()
	if err != nil {
		return err
	}
	for idx, link := range links {
		if link.ID == linkID {
			linkIndex = idx

			break
		}
	}
	if linkIndex == -1 {
		return c.SendStatus(http.StatusNotFound)
	}
	err = writeLinksToFile(append(links[:linkIndex], links[linkIndex+1:]...))
	if err != nil {
		return err
	}

	return c.SendStatus(http.StatusNoContent)
}

func (p *API) HandleLinks(r fiber.Router) {
	r.Get("", getLinksHandler)
	r.Post("", addLinkHandler)
	r.Delete("/:id", deleteLinkHandler)
}
