package user

import (
	"context"

	"github.com/artHouse-Docs/backend/pkg/config"
	db "github.com/artHouse-Docs/backend/pkg/database"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserCollenction() (coll *mongo.Collection, err error) {
	cfg := config.Configure()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	client, err := db.NewClient(ctx, cfg.Database.Host, cfg.Database.Port, cfg.Database.Username, cfg.Database.Password)
	if err != nil {
		return nil, err
	}
	return db.NewCollection(ctx, client, "users"), nil
}
