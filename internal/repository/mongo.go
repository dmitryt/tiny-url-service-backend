package repository

import (
	"context"
	"errors"
	"fmt"

	// is used for init postgres.
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ErrDBOpen = errors.New("database open error")

type MongoRepo struct {
	client *mongo.Client
}

func (r *MongoRepo) Connect(ctx context.Context, url string) (err error) {
	log.Debug().Msgf("Connecting to %s", url)
	r.client, err = mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		return fmt.Errorf("%s: %w", ErrDBOpen, err)
	}

	return
}

func (r *MongoRepo) Close(ctx context.Context) error {
	return r.client.Disconnect(ctx)
}

func NewMongoRepo() *MongoRepo {
	return &MongoRepo{}
}
