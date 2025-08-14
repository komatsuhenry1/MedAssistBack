package repository

import (
	"context"
	"errors"
	"fmt"
	"medassist/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	FindUserByEmail(email string) (model.User, error)
	CreateUser(user *model.User) error
}

type userRepository struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &userRepository{
		collection: db.Collection("users"),
		ctx:        context.Background(),
	}
}

func (r *userRepository) FindUserByEmail(email string) (model.User, error) {

	var user model.User
	err := r.collection.FindOne(r.ctx, bson.M{"email": email}).Decode(&user)
	fmt.Println(err)
	if err == nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return user, fmt.Errorf("usuário não encontrado")
		}
		return user, err
	}

	return user, nil
}

func (r *userRepository) CreateUser(user *model.User) error {
	_, err := r.collection.InsertOne(r.ctx, user)
	return err
}
