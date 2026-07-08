package auth

import "errors"

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrRepository         = errors.New("repository error")
	ErrInvalidCredentials = errors.New("invalid credentials")
)
