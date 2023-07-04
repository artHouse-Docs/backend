package user

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/artHouse-Docs/backend/pkg/hashing"
)

// handle request() {userId     register()}

func (u *User) Register(ctx context.Context) (err error) {
	coll, err := NewUserCollenction()
	if err != nil {
		return errors.New("database unavaliable")
	}

	result, err := coll.InsertOne(ctx, bson.D{
		{"name", u.Name},
		{"surname", u.Surname},
		{"password", hashing.MakeHash(u.PasswordHash)},
		{"email", u.Email},
	})

	if err != nil {
		return errors.New("server error")
	}

	u.ID = result.InsertedID.(primitive.ObjectID).Hex()

	return nil
}

func (u *User) Login(ctx context.Context) (result bool, err error) {
	coll, err := NewUserCollenction()
	if err != nil {
		return false, errors.New("database unavaliable")
	}

	var userByEmail User

	coll.FindOne(ctx, bson.D{
		{"email", u.Email},
	}).Decode(&userByEmail)

	if userByEmail.ID == "" {
		return false, errors.New("user not found")
	} else if !hashing.CompareHash(u.PasswordHash, userByEmail.PasswordHash) {
		return false, errors.New("wrong password")
	} else {
		u.ID = userByEmail.ID
		u.Name = userByEmail.Name
		u.Surname = userByEmail.Surname
	}
	return true, nil
}
