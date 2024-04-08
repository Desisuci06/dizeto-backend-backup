package utils

import "errors"

var (
	ErrInvalidCredentials      = errors.New("invalid credentials")
	ErrIsUsernameOrEmailExists = errors.New("username or email already exists")
)
