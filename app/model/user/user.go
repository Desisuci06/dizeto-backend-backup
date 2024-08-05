package user

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// Struct User dengan validasi
type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	Name      string    `gorm:"unique;not null" json:"name" validate:"required"`
	Email     string    `gorm:"unique;not null" json:"email" validate:"required,email"`
	Username  string    `gorm:"unique;not null" json:"username" validate:"required,min=3,max=50"`
	Password  string    `gorm:"not null" json:"password" validate:"required,min=6"`
	Image     string    `json:"image"`
	Role      string    `json:"role"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// Validate melakukan validasi pada struct User
func (u *User) Validate() error {
	validate := validator.New()
	err := validate.Struct(u)
	if err != nil {
		return err
	}

	return nil
}
