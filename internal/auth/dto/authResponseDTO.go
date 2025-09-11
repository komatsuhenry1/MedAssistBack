package dto

import "go.mongodb.org/mongo-driver/bson/primitive"


type AuthUser struct {
    ID       primitive.ObjectID `bson:"_id"` 
	Name     string `bson:"name"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
	Role     string `bson:"role"`
	Hidden   bool   `bson:"hidden"`
	TempCode int    `bson:"temp_code"`
}

type CodeResponseDTO struct {
	Code int `json:"code"`
}
