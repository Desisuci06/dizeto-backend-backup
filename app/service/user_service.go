package service

import (
	model "dizeto-backend/app/model/user"

	"dizeto-backend/app/repository"
	"dizeto-backend/utils"

	"github.com/google/uuid"
)

type UserService interface {
	Register(username, password, email, name, image string) error
	Login(username, password string) (*model.User, string, error)
	IsUsernameOrEmailExists(username, email string) bool
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo}
}

func (us *userService) Register(username, password, email, name, image string) error {
	// Generate UUID for user ID
	userID, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	if us.IsUsernameOrEmailExists(username, email) {
		return utils.ErrIsUsernameOrEmailExists
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	// Create new user
	newUser := &model.User{
		ID:       userID,
		Name:     name,
		Image:    image,
		Username: username,
		Password: hashedPassword,
		Email:    email,
		Role:     "customer",
	}
	err = us.userRepo.CreateUser(newUser)
	if err != nil {
		return err
	}

	return nil
}

func (us *userService) Login(username, password string) (*model.User, string, error) {
	// Retrieve user by username
	user, err := us.userRepo.GetUserByUsername(username)
	if err != nil {
		return nil, "", err
	}

	// Check password
	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, "", utils.ErrInvalidCredentials
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(username, password, user.Role)
	if err != nil {
		return nil, "", err
	}
	return user, token, nil
}

func (us *userService) IsUsernameOrEmailExists(username, email string) bool {
	// Check if username or email already exists in the database
	userByUsername, err := us.userRepo.GetUserByUsername(username)
	if err == nil && userByUsername != nil {
		return true
	}

	userByEmail, err := us.userRepo.GetUserByEmail(email)
	if err == nil && userByEmail != nil {
		return true
	}

	return false
}
