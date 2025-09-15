package repository

import (
	"context"
	"errors"
	"fmt"
	"medassist/internal/auth/dto"
	"medassist/internal/model"
	"medassist/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

type UserRepository interface {
	FindUserByEmail(email string) (dto.AuthUser, error)
	FindUserByCpf(cpf string) (model.User, error)
	FindUserById(id string) (model.User, error)
	CreateUser(user *model.User) error
	UpdateTempCode(userID string, code int) error
	UpdateUser(userId string, userUpdated bson.M) (model.User, error)
	UpdateUserFields(userId string, updates map[string]interface{}) (model.User, error)
	UserExistsByEmail(email string) (bool, error)
	FindAuthUserByID(id string) (dto.AuthUser, error)
	UpdatePasswordByUserID(userID string, hashedPassword string) error
	DownloadFileByID(fileID primitive.ObjectID) (*gridfs.DownloadStream, error)
}

type userRepository struct {
	collection *mongo.Collection
	ctx        context.Context
	bucket     *gridfs.Bucket
}

func NewUserRepository(db *mongo.Database) UserRepository {
	bucket, err := gridfs.NewBucket(db)
	if err != nil {
		panic(err)
	}
	return &userRepository{
		collection: db.Collection("users"),
		ctx:        context.Background(),
		bucket:     bucket,
	}
}

func (r *userRepository) FindUserByEmail(email string) (dto.AuthUser, error) {
	var authUser dto.AuthUser

	err := r.collection.FindOne(r.ctx, bson.M{"email": email}).Decode(&authUser)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return authUser, fmt.Errorf("usuário não encontrado")
		}
		return authUser, err
	}

	return authUser, nil
}

func (r *userRepository) FindAuthUserByID(id string) (dto.AuthUser, error) {
	var authUser dto.AuthUser

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return authUser, fmt.Errorf("ID inválido")
	}

	err = r.collection.FindOne(r.ctx, bson.M{"_id": objectID}).Decode(&authUser)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return authUser, fmt.Errorf("usuário não encontrado")
		}
		return authUser, err
	}
	return authUser, nil

}

func (r *userRepository) UpdatePasswordByUserID(userID string, hashedPassword string) error {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return fmt.Errorf("ID inválido")
	}

	result, err := r.collection.UpdateByID(r.ctx, objID, bson.M{
		"$set": bson.M{
			"password":   hashedPassword,
			"updated_at": time.Now(),
		},
	})
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return fmt.Errorf("nenhum usuário encontrado com o ID %s", userID)
	}
	return nil
}

func (r *userRepository) FindUserByCpf(cpf string) (model.User, error) {

	var user model.User
	err := r.collection.FindOne(r.ctx, bson.M{"cpf": cpf}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return user, fmt.Errorf("usuário não encontrado")
		}
		return user, err
	}

	return user, nil
}

func (r *userRepository) FindUserById(id string) (model.User, error) {
	var user model.User

	// converter para ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, fmt.Errorf("ID inválido: %w", err)
	}

	err = r.collection.FindOne(r.ctx, bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return model.User{}, fmt.Errorf("usuário não encontrado")
		}
		return model.User{}, err
	}

	return user, nil
}

func (r *userRepository) CreateUser(user *model.User) error {
	_, err := r.collection.InsertOne(r.ctx, user)
	return err
}

func (r *userRepository) UpdateTempCode(userID string, code int) error {

	// converter para ObjectID
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

func (r *userRepository) UpdateUser(userId string, userUpdates bson.M) (model.User, error) {
	if titleRaw, ok := userUpdates["title"]; ok {
		title, ok := titleRaw.(string)
		if ok {
			formattedTitle := utils.CapitalizeFirstWord(title)
			userUpdates["name"] = formattedTitle
		}
	}

	product, err := r.UpdateUserFields(userId, userUpdates)
	if err != nil {
		return model.User{}, fmt.Errorf("erro ao atualizar produto")
	}
	return product, nil
}

func (r *userRepository) UpdateUserFields(id string, updates map[string]interface{}) (model.User, error) {
	cleanUpdates := bson.M{}

	for key, value := range updates {
		if value != nil {
			cleanUpdates[key] = value
		}
	}

	if len(cleanUpdates) == 0 {
		return model.User{}, fmt.Errorf("nenhum campo válido para atualizar")
	}

	cleanUpdates["updated_at"] = time.Now()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.User{}, fmt.Errorf("ID inválido")
	}

	update := bson.M{"$set": cleanUpdates}

	_, err = r.collection.UpdateByID(context.TODO(), objID, update)
	if err != nil {
		return model.User{}, err
	}

	return r.FindUserById(id)
}

func (r *userRepository) UserExistsByEmail(email string) (bool, error) {
	err := r.collection.FindOne(r.ctx, bson.M{"email": email}).Err()
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r *userRepository) DownloadFileByID(fileID primitive.ObjectID) (*gridfs.DownloadStream, error) {
	// Usa o bucket do GridFS para abrir o stream de download.
	downloadStream, err := r.bucket.OpenDownloadStream(fileID)
	if err != nil {
		// Este erro ocorrerá se o fileID não existir no GridFS.
		return nil, err
	}

	return downloadStream, nil
}
