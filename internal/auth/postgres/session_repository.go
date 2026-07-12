package postgres

import (
	"context"
	"fmt"
	"errors"
	"cloud-storage/internal/auth"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/google/uuid"
)

type SessionRepository struct {
    db *pgxpool.Pool
}

var _ auth.SessionRepository = (*SessionRepository)(nil)

func NewSessionRepository(db *pgxpool.Pool) *SessionRepository {
    return &SessionRepository {
        db: db,
    }
}

func (r *SessionRepository) Create(ctx context.Context, session *auth.Session) error {
	query := `
        INSERT INTO sessions (id, user_id, token_hash, expires_at, user_agent, ip_address)
        VALUES ($1, $2, $3, $4, $5, $6)
    `
	_, err := r.db.Exec(
		ctx,
		query,
		session.SessionID,
		session.UserID,
		session.TokenHash,
		session.ExpiresAt,
		session.UserAgent,
		session.IPAddress,
	)
	if err != nil {
		return fmt.Errorf("%w: create session: %w", auth.ErrRepository, err)
	}
	return nil
}
	
func (r *SessionRepository) FindByID(ctx context.Context, id uuid.UUID) (*auth.Session, error) {
	query := `
		SELECT 
			id, user_id, token_hash, expires_at, created_at, 
			last_used_at, revoked, user_agent, ip_address
		FROM sessions
		WHERE id = $1	
	`
	var session auth.Session

	err := r.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&session.SessionID,
		&session.UserID,
		&session.TokenHash,
		&session.ExpiresAt,
		&session.CreatedAt,
		&session.LastUsedAt,
		&session.Revoked,
		&session.UserAgent,
		&session.IPAddress,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) { return nil, auth.ErrSessionNotFound }
		return nil, fmt.Errorf("%w: find session by id: %w", auth.ErrRepository, err)
	}
	return &session, nil 
}

func (r *SessionRepository) FindByTokenHash(ctx context.Context, hash string) (*auth.Session, error) {
	query := `
		SELECT 
			id, user_id, token_hash, expires_at, created_at, 
			last_used_at, revoked, user_agent, ip_address
		FROM sessions
		WHERE token_hash = $1	
	`
	var session auth.Session

	err := r.db.QueryRow(
		ctx,
		query,
		hash,
	).Scan(
		&session.SessionID,
		&session.UserID,
		&session.TokenHash,
		&session.ExpiresAt,
		&session.CreatedAt,
		&session.LastUsedAt,
		&session.Revoked,
		&session.UserAgent,
		&session.IPAddress,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) { return nil, auth.ErrSessionNotFound }
		return nil, fmt.Errorf("%w: find session by token hash: %w", auth.ErrRepository, err)
	}
	return &session, nil 
}
	
func (r *SessionRepository) FindByUserID(ctx context.Context, id uint) ([]*auth.Session, error) {
	query := `
		SELECT 
			id, user_id, token_hash, expires_at, created_at, 
			last_used_at, revoked, user_agent, ip_address
		FROM sessions
		WHERE user_id = $1	
	`
	rows, err := r.db.Query(
		ctx,
		query,
		id,
	)
	if err != nil {
		return nil, fmt.Errorf("%w: find sessions by user id: %w", auth.ErrRepository, err)
	}
	defer rows.Close()

	sessions := make([]*auth.Session, 0)

	for rows.Next() {
		var session auth.Session

		err := rows.Scan(
			&session.SessionID,
			&session.UserID,
			&session.TokenHash,
			&session.ExpiresAt,
			&session.CreatedAt,
			&session.LastUsedAt,
			&session.Revoked,
			&session.UserAgent,
			&session.IPAddress,
		)
		if err != nil {
			return nil, fmt.Errorf("%w: scan session: %w", auth.ErrRepository, err)
		}

		sessions = append(sessions, &session)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%w: iterate sessions: %w", auth.ErrRepository, err)
	}
	return sessions, nil
}
	
func (r *SessionRepository) UpdateLastUsedAt(ctx context.Context, id uuid.UUID) error {
	query := `
		UPDATE sessions
		SET last_used_at = now()
		WHERE id = $1
	`
	result, err := r.db.Exec(
		ctx,
		query,
		id,
	)
	if err != nil {
		return fmt.Errorf("%w: update last used at: %w", auth.ErrRepository, err)
	}
	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return auth.ErrSessionNotFound
	}
	return nil
}
	
func (r *SessionRepository) Revoke(ctx context.Context, id uuid.UUID) error {
	query := `
		UPDATE sessions
		SET revoked = true
		WHERE id = $1
	`
	result, err := r.db.Exec(
		ctx,
		query,
		id,
	)
	if err != nil {
		return fmt.Errorf("%w: revoke session: %w", auth.ErrRepository, err)
	}
	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return auth.ErrSessionNotFound
	}
	return nil

}
	
func (r *SessionRepository) RevokeAllByUserID(ctx context.Context, id uint) error {
	query := `
		UPDATE sessions
		SET revoked = true
		WHERE user_id = $1
	`
	result, err := r.db.Exec(
		ctx,
		query,
		id,
	)
	if err != nil {
		return fmt.Errorf("%w: revoke user sessions: %w", auth.ErrRepository, err)
	}
	if result.RowsAffected() == 0 {
		return auth.ErrSessionNotFound
	}
	return nil
}
	
func (r *SessionRepository) DeleteExpired(ctx context.Context) error {
	query := `
		DELETE FROM sessions
		WHERE expires_at <= now()
	`
	_, err := r.db.Exec(
		ctx,
		query,
	)
	if err != nil {
		return fmt.Errorf("%w: delete expired sessions: %w", auth.ErrRepository, err)
	}
	return nil
}
