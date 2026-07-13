package postgres

import (
	"cloud-storage/internal/auth/entity"
	"cloud-storage/internal/auth/repository"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
    db *pgxpool.Pool
}

var _ repository.UserRepository = (*UserRepository)(nil)

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
    return &UserRepository {
        db: db,
    }
}

func (r *UserRepository) Create(ctx context.Context, user *entity.User) error {
	query := `
        INSERT INTO users (name, email, password_hash)
        VALUES ($1, $2, $3)
        RETURNING id
    `

	err := r.db.QueryRow(
        ctx,
        query,
        user.Name,
        user.Email,
        user.PasswordHash,
    ).Scan(&user.ID)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return repository.ErrEmailAlreadyExists
			}
		}
		return fmt.Errorf("%w: create user: %w", repository.ErrRepository, err)
	}
	return nil
}

func (r *UserRepository) FindByID(ctx context.Context, id uint) (*entity.User, error) {
	var user entity.User

    err := r.db.QueryRow(
        ctx,
        `
        SELECT id,name,email,password_hash, created_at
        FROM users
        WHERE id=$1
        `,
        id,
    ).Scan(
        &user.ID,
        &user.Name,
        &user.Email,
        &user.PasswordHash,
		&user.CreatedAt,
    )

   	if errors.Is(err, pgx.ErrNoRows) {
		return nil, repository.ErrUserNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("%w: find user by id: %w", repository.ErrRepository, err)
	}
	return &user, nil 
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User

    err := r.db.QueryRow(
        ctx,
        `
        SELECT id,name,email,password_hash, created_at
        FROM users
        WHERE email=$1
        `,
        email,
    ).Scan(
        &user.ID,
        &user.Name,
        &user.Email,
        &user.PasswordHash,
		&user.CreatedAt,
    )

  	if errors.Is(err, pgx.ErrNoRows) {
		return nil, repository.ErrUserNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("%w: find user by email: %w", repository.ErrRepository, err)
	}
	return &user, nil 
}
