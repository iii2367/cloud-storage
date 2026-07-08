package postgres

import (
	"context"
	"fmt"
	"errors"
	"cloud-storage/internal/auth"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
    db *pgxpool.Pool
}

var _ auth.Repository = (*Repository)(nil)

func New(db *pgxpool.Pool) *Repository {
    return &Repository {
        db: db,
    }
}

func (r *Repository) Create(ctx context.Context, user *auth.User) error {
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
				return auth.ErrEmailAlreadyExists
			}
		}
		return fmt.Errorf("%w: create user: %w", auth.ErrRepository, err)
	}
	return nil
}

func (r *Repository) FindByID(ctx context.Context, id uint) (*auth.User, error) {
	var user auth.User

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
		return nil, auth.ErrUserNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("%w: find user by id: %w", auth.ErrRepository, err)
	}
	return &user, nil 
}

func (r *Repository) FindByEmail(ctx context.Context, email string) (*auth.User, error) {
	var user auth.User

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
		return nil, auth.ErrUserNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("%w: find user by email: %w", auth.ErrRepository, err)
	}
	return &user, nil 
}
