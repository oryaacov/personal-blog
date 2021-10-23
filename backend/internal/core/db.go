package core

import (
	"context"
	"fmt"
	"time"

	"github.com/oryaacov/personal-blog/internal/common"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DB struct {
	*mongo.Client
	deps Deps
}

type Deps struct {
	Config common.Config
}

func InitDB(config common.Config) (*DB, error) {
	deps := Deps{Config: config}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://%s", deps.Config.Database.Address)))
	if err != nil {
		return nil, fmt.Errorf("failed to open DB connection with: %w", err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, fmt.Errorf("failed to ping DB connection with: %w", err)
	}

	return &DB{Client: client, deps: deps}, nil
}

func (c *DB) GetCollection(name string) *mongo.Collection {
	return c.Database(c.deps.Config.Database.Name).Collection(name)
}
