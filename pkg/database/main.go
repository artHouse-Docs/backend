package database

import (
	"context"

	"github.com/artHouse-Docs/backend/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(ctx context.Context, host, port, username, password string) (*mongo.Database, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+username+":"+password+"@"+host+":"+port))
	if err != nil {
		return nil, err
	}
	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}
	return client.Database("main"), err
}

func NewCollection(ctx context.Context, name string) (coll *mongo.Collection, err error) {
	cfg := config.Configure()

	client, err := NewClient(ctx, cfg.Database.Host, cfg.Database.Port, cfg.Database.Username, cfg.Database.Password)
	if err != nil {
		return nil, err
	}

	db := client.Client().Database("main")

	return db.Collection(name), nil
}
