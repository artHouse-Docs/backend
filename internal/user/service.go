package user

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/artHouse-Docs/backend/pkg/hashing"
)

// handle request() {userId     register()}

func (u *UserDTO) Register(ctx context.Context) (err error) {
	coll, err := NewUserCollenction()
	if err != nil {
		return errors.New("database unavaliable")
	}

	result, err := coll.InsertOne(ctx, bson.D{
		{"name", u.Name},
		{"surname", u.Surname},
		{"password", hashing.MakeHash(u.Password)},
		{"email", u.Email},
	})

	if err != nil {
		return errors.New("server error")
	}

	u.ID = result.InsertedID.(primitive.ObjectID).Hex()

	return nil
}
