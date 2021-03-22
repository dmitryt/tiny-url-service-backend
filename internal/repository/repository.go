package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/dmitryt/tiny-url-service-backend/internal/config"
	models "github.com/dmitryt/tiny-url-service-backend/internal/models"
)

var ErrUnSupportedRepoType = errors.New("unsupported repository type")

type dbConnector interface {
	Connect(context.Context, string) error
	Close() error
}

type CRUD interface {
	GetLinks() ([]models.Link, error)

	CreateLink(models.Link) (models.Link, error)
	DeleteEvent(int64) error

	CreateUser(models.User) (models.User, error)
	GetUser(int64) (models.User, error)
	dbConnector
}

func newRepo(repoType string) interface{} {
	if repoType == "mongo" {
		return NewMongoRepo()
	}

	return nil
}

func NewCRUD(repoType string) CRUD {
	repo, ok := newRepo(repoType).(CRUD)
	if !ok {
		return nil
	}

	return repo
}

func GetMongoURI(c *config.DBConfig) string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", c.User, c.Password, c.Host, c.Port, c.DBName)
}
