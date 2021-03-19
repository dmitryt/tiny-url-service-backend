package repository

import (
	"context"
	"errors"
	"fmt"
	"time"
)

var ErrUnSupportedRepoType = errors.New("unsupported repository type")

type dbConnector interface {
	Connect(context.Context, string) error
	Close() error
}

type CRUD interface {
	GetLinks() ([]Link, error)

	CreateLink(Link) (Link, error)
	DeleteEvent(int64) error

	CreateUser(User) (User, error)
	GetUser(int64) (User, error)
	dbConnector
}

func newRepo(repoType string, args ...interface{}) interface{} {
	switch repoType {
	case "mongo":
		return NewMongoRepo()
	}
	return nil
}

func NewCRUD(repoType string, args ...interface{}) CRUD {
	repo, ok := newRepo(repoType, args...).(CRUD)
	if !ok {
		return nil
	}
	return repo
}

func GetMongoURI(c *config.DBConfig) string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", c.User, c.Password, c.Host, c.Port, c.DBName)
}
