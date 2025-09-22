package repository

import (
	"context"
	"errors"
	"fmt"
	"io"
	"medassist/internal/auth/dto"
	"medassist/internal/model"
	"medassist/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

type NurseRepository interface {
	FindNurseByEmail(email string) (dto.AuthUser, error)
	FindNurseByCpf(cpf string) (model.Nurse, error)
	FindNurseById(id string) (model.Nurse, error)
	CreateNurse(nurse *model.Nurse) error
	FindAllNurses() ([]model.Nurse, error)
	FindAllNursesNotVerified() ([]model.Nurse, error)
	UpdateTempCode(userID string, code int) error
	UpdateNurse(nurseId string, userUpdated bson.M) (model.Nurse, error)
	UpdateNurseFields(userId string, updates map[string]interface{}) (model.Nurse, error)
	SetLicenseDocumentID(nurseID, documentID primitive.ObjectID) error
	UploadFile(file io.Reader, fileName string) (primitive.ObjectID, error)
	FindAuthNurseByID(id string) (dto.AuthUser, error)
	UpdatePasswordByNurseID(userID string, hashedPassword string) error
	GetIdsNursesPendents() ([]string, error)
}

type nurseRepository struct {
	collection *mongo.Collection
	ctx        context.Context
	bucket     *gridfs.Bucket
}

func NewNurseRepository(db *mongo.Database) NurseRepository {
	bucket, err := gridfs.NewBucket(db)
	if err != nil {
		panic(err)
	}

	return &nurseRepository{
		collection: db.Collection("nurses"),
		ctx:        context.Background(),
		bucket:     bucket,
	}
}

func (r *nurseRepository) FindNurseByEmail(email string) (dto.AuthUser, error) {
	var authUser dto.AuthUser

	// A busca é feita na coleção "nurses"
	err := r.collection.FindOne(r.ctx, bson.M{"email": email}).Decode(&authUser)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return authUser, fmt.Errorf("usuário não encontrado")
		}
		return authUser, err
	}

	return authUser, nil
}

func (r *nurseRepository) FindAuthNurseByID(id string) (dto.AuthUser, error) {
	var authUser dto.AuthUser

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return authUser, fmt.Errorf("ID inválido")
	}

    err = r.collection.FindOne(r.ctx, bson.M{"_id": objectID}).Decode(&authUser)
    if err != nil {
        if errors.Is(err, mongo.ErrNoDocuments) {
            return authUser, fmt.Errorf("enfermeiro não encontrado")
        }
        return authUser, err
    }
    return authUser, nil
}

func (r *nurseRepository) UpdatePasswordByNurseID(userID string, hashedPassword string) error {
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

func (r *nurseRepository) FindNurseByCpf(cpf string) (model.Nurse, error) {

	var nurse model.Nurse
	err := r.collection.FindOne(r.ctx, bson.M{"cpf": cpf}).Decode(&nurse)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nurse, fmt.Errorf("enfermeiro(a) não encontrado")
		}
		return nurse, err
	}

	return nurse, nil
}

func (r *nurseRepository) FindNurseById(id string) (model.Nurse, error) {
	var nurse model.Nurse

	// converter para ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nurse, fmt.Errorf("ID inválido: %w", err)
	}

	err = r.collection.FindOne(r.ctx, bson.M{"_id": objectID}).Decode(&nurse)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return model.Nurse{}, fmt.Errorf("enfermeiro(a) não encontrado(a)")
		}
		return model.Nurse{}, err
	}

	return nurse, nil
}

func (r *nurseRepository) CreateNurse(nurse *model.Nurse) error {
	_, err := r.collection.InsertOne(r.ctx, nurse)
	return err
}

func (r *nurseRepository) UploadFile(file io.Reader, fileName string) (primitive.ObjectID, error) { // retorna o object id que foi criado em fs.files
	uploadStream, err := r.bucket.OpenUploadStream(fileName)
	if err != nil {
		return primitive.NilObjectID, err
	}
	defer uploadStream.Close()

	if _, err := io.Copy(uploadStream, file); err != nil {
		return primitive.NilObjectID, err
	}

	fileID := uploadStream.FileID.(primitive.ObjectID)
	return fileID, nil
}

func (r *nurseRepository) SetLicenseDocumentID(nurseID, documentID primitive.ObjectID) error {
	filter := bson.M{"_id": nurseID}
	update := bson.M{"$set": bson.M{"license_document_id": documentID, "updated_at": time.Now()}}
	_, err := r.collection.UpdateOne(r.ctx, filter, update)
	return err
}

func (r *nurseRepository) UpdateTempCode(userID string, code int) error {

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

func (r *nurseRepository) UpdateNurse(nurseId string, nurseUpdates bson.M) (model.Nurse, error) {
	if titleRaw, ok := nurseUpdates["title"]; ok {
		title, ok := titleRaw.(string)
		if ok {
			formattedTitle := utils.CapitalizeFirstWord(title)
			nurseUpdates["name"] = formattedTitle
		}
	}

	nurse, err := r.UpdateNurseFields(nurseId, nurseUpdates)
	if err != nil {
		return model.Nurse{}, fmt.Errorf("erro ao atualizar enfermeiro(a)")
	}
	return nurse, nil
}

func (r *nurseRepository) UpdateNurseFields(id string, updates map[string]interface{}) (model.Nurse, error) {
	cleanUpdates := bson.M{}

	for key, value := range updates {
		if value != nil {
			cleanUpdates[key] = value
		}
	}

	if len(cleanUpdates) == 0 {
		return model.Nurse{}, fmt.Errorf("nenhum campo válido para atualizar")
	}

	cleanUpdates["updated_at"] = time.Now()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Nurse{}, fmt.Errorf("ID inválido")
	}

	update := bson.M{"$set": cleanUpdates}

	_, err = r.collection.UpdateByID(context.TODO(), objID, update)
	if err != nil {
		return model.Nurse{}, err
	}

	return r.FindNurseById(id)
}

func (r *nurseRepository) GetIdsNursesPendents() ([]string, error){
	var nursesIds []string
	filter := bson.M{"verification_seal": false}
	cursor, err := r.collection.Find(r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(r.ctx)

	for cursor.Next(r.ctx) {
		var nurse model.Nurse
		if err := cursor.Decode(&nurse); err != nil {
			return nil, err
		}
		nursesIds = append(nursesIds, nurse.ID.Hex())
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return nursesIds, nil
}

func (r *nurseRepository) FindAllNurses() ([]model.Nurse, error) {
	var nurses []model.Nurse

	cursor, err := r.collection.Find(r.ctx, bson.M{})
	if err != nil {
		return nurses, err
	}
	defer cursor.Close(r.ctx)

	if err = cursor.All(r.ctx, &nurses); err != nil {
		return nurses, err
	}

	return nurses, nil
}

func (r *nurseRepository) FindAllNursesNotVerified() ([]model.Nurse, error) {
	var nurses []model.Nurse

	filter := bson.M{
		"verification_seal": false,
	}

	cursor, err := r.collection.Find(r.ctx, filter)
	if err != nil {
		return nurses, err
	}
	defer cursor.Close(r.ctx)

	if err = cursor.All(r.ctx, &nurses); err != nil {
		return nurses, err
	}

	return nurses, nil
}