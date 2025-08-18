package repository

import (
	"context"
	"errors"
	"fmt"
	"medassist/internal/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	FindUserByEmail(email string) (model.User, error)
	CreateUser(user *model.User) error
	UpdateTempCode(userID string, code int) error
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
	fmt.Println("err", err)
	if err != nil {
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

func (r *userRepository) UpdateTempCode(userID string, code int) error {
	fmt.Println("userID: ", userID)
	fmt.Println("code: ", code)

	// Converter para ObjectID
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return fmt.Errorf("ID inválido: %w", err)
	}

	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"temp_code": code,
			"updatedAt": time.Now(),
		},
	}

	result, err := r.collection.UpdateOne(r.ctx, filter, update)
	if err != nil {
		return fmt.Errorf("erro ao atualizar temp_code: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("nenhum documento encontrado com o ID informado")
	}

	return nil
}
