package repository

import ("errors")

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrEmailAlreadyExists = errors.New("email already exists")

	ErrSessionNotFound	  = errors.New("session not found")

	ErrRepository         = errors.New("repository error")
)
