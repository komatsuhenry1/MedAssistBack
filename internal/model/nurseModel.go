package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Nurse struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name     string             `bson:"name" json:"name" binding:"required"`
	Email    string             `bson:"email" json:"email" binding:"required,email"`
	Phone    string             `bson:"phone" json:"phone" binding:"required,phone"`
	Address  string             `bson:"address" json:"address" binding:"required,address"`
	Cpf      string             `bson:"cpf" json:"cpf" binding:"required"`
	PixKey   string             `bson:"pix_key" json:"pix_key" binding:"required"`
	Password string             `bson:"password" json:"password" binding:"required"`

	LicenseNumber   string `bson:"license_number" json:"license_number" binding:"required"` // registro profissional
	Specialization  string `bson:"specialization" json:"specialization"`                    // área (ex: pediatrics, geriatrics, ER)
	Shift           string `bson:"shift" json:"shift"`                                      // manhã, tarde, noite
	Department      string `bson:"department" json:"department"`                            // setor/hospital onde trabalha
	YearsExperience int    `bson:"years_experience" json:"years_experience"`
	// Available       bool               `bson:"available" json:"available"` // disponível para consultas

	Hidden      bool      `bson:"hidden" json:"hidden"`
	Role        string    `bson:"role" json:"role" binding:"required"`
	Online      bool      `bson:"online" json:"online" binding:"required"`
	FirstAccess bool      `bson:"first_access" json:"first_access"`
	CreatedAt   time.Time `bson:"created_at" json:"created_at"`
	TempCode    int       `bson:"temp_code" json:"temp_code"`
	UpdatedAt   time.Time `bson:"updated_at" json:"updated_at"`
}
