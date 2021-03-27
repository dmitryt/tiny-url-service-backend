package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"

	models "github.com/dmitryt/tiny-url-service-backend/internal/models"
	"github.com/dmitryt/tiny-url-service-backend/internal/session"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LinkPayload struct {
	Value string `json:"full_url"`
}

var linksStoragePath = filepath.Join("fixtures", "links.json")

func logErrorIfExists(ctxErr, err error) error {
	if err != nil {
		log.Error().Msgf("%s: %s", ctxErr, err)

		return ctxErr
	}

	return nil
}

func getCurrentUserID(c *fiber.Ctx) (string, error) {
	store, err := session.Store.Get(c)
	if err != nil {
		return "", logErrorIfExists(fiber.ErrUnauthorized, err)
	}
	uid := fmt.Sprintf("%v", store.Get("uid"))
	if uid == "" {
		return "", logErrorIfExists(fiber.ErrUnauthorized, err)
	}

	return uid, nil
}

func readLinksFromFile(c *fiber.Ctx) (result []models.Link, err error) {
	var allData []models.Link
	uid, err := getCurrentUserID(c)
	if err != nil {
		return
	}
	content, err := ioutil.ReadFile(linksStoragePath)
	if err != nil {
		return
	}
	if len(content) == 0 {
		content = []byte("[]")
	}

	err = json.Unmarshal(content, &allData)

	if err != nil {
		return result, logErrorIfExists(ErrContentUnmarshal, err)
	}

	for _, link := range allData {
		if link.User == uid {
			result = append(result, link)
		}
	}

	return result, nil
}

func writeLinksToFile(data []models.Link) (err error) {
	content, err := json.Marshal(data)
	if err != nil {
		return logErrorIfExists(ErrContentMarshal, err)
	}
	err = ioutil.WriteFile(linksStoragePath, content, 0o600)

	return logErrorIfExists(ErrFileWrite, err)
}

func getLinksHandler(c *fiber.Ctx) error {
	links, err := readLinksFromFile(c)
	if err != nil {
		return err
	}

	return c.JSON(links)
}

func addLinkHandler(c *fiber.Ctx) error {
	uid, err := getCurrentUserID(c)
	if err != nil {
		return err
	}

	payload := new(LinkPayload)
	if err := c.BodyParser(payload); err != nil {
		return logErrorIfExists(ErrParsingBody, err)
	}
	if payload.Value == "" {
		return logErrorIfExists(ErrValidationAddItem, errors.New("invalid data"))
	}
	links, err := readLinksFromFile(c)
	if err != nil {
		return err
	}
	newLink := models.Link{ID: primitive.NewObjectID().Hex(), Alias: fmt.Sprintf("/%s", uuid.New().String()), URL: payload.Value, User: uid}
	err = writeLinksToFile(append(links, newLink))
	if err != nil {
		return err
	}

	return c.JSON(newLink)
}

func deleteLinkHandler(c *fiber.Ctx) error {
	linkID := c.Params("id")
	linkIndex := -1
	links, err := readLinksFromFile(c)
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
