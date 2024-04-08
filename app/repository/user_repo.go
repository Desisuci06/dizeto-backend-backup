package repository

import (
	model "dizeto-backend/app/model/user"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	CreateUser(user *model.User) error
	GetUserByUsername(username string) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
}

// Implementasi repository menggunakan Gorm
type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (ur *userRepository) CreateUser(user *model.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	err := ur.db.Create(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (ur *userRepository) GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	err := ur.db.Where("username = ?", username).First(&user).Error
	return &user, err
}

func (ur *userRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := ur.db.Where("email = ?", email).First(&user).Error
	return &user, err
}
