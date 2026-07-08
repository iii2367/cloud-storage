package auth

import ("context")

type Repository interface {
	Create(ctx context.Context, user *User) error
    FindByEmail(ctx context.Context, email string) (*User, error)
    FindByID(ctx context.Context, id uint) (*User, error)
}
