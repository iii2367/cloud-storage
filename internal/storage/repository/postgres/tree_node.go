package postgres

import (
	"context"
	"errors"
	"fmt"

	"cloud-storage/internal/storage/entity"
	"cloud-storage/internal/storage/repository"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TreeNodeRepository struct {
	db *pgxpool.Pool
}

var _ repository.TreeNodeRepository = (*TreeNodeRepository)(nil)

func NewTreeNodeRepository(db *pgxpool.Pool) *TreeNodeRepository {
	return &TreeNodeRepository{
		db: db,
	}
}

func (r *TreeNodeRepository) Create(ctx context.Context, node *entity.TreeNode) error {
	query := `
		INSERT INTO tree_nodes (id, parent_id, user_id, file_type, extension, mime_type, name, description, size)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
	`
	_, err := r.db.Exec(
		ctx,
		query,
		node.ID,
		node.ParentID,
		node.UserID,
		node.FileType,
		node.Extension,
		node.MimeType,
		node.Name,
		node.Description,
		node.Size,
	)
	if err != nil {
		return fmt.Errorf("%w: create tree node: %w", repository.ErrRepository, err)
	}
	return nil
}

func (r *TreeNodeRepository) FindByID(ctx context.Context, id uuid.UUID, userID uint) (*entity.TreeNode, error) {
	query := `
	SELECT
		id, parent_id, user_id, file_type, extension, mime_type,
		upload_at, updated_at, name, description, size
		FROM tree_nodes
		WHERE id = $1 AND user_id = $2
	`
	var node entity.TreeNode

	err := r.db.QueryRow(
		ctx,
		query,
		id,
		userID,
	).Scan(
		&node.ID,
		&node.ParentID,
		&node.UserID,
		&node.FileType,
		&node.Extension,
		&node.MimeType,
		&node.UploadAt,
		&node.UpdatedAt,
		&node.Name,
		&node.Description,
		&node.Size,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, repository.ErrTreeNodeNotFound
		}
		return nil, fmt.Errorf("%w: find tree node by id: %w", repository.ErrRepository, err)
	}
	return &node, nil
}

func (r *TreeNodeRepository) FindByUserID(ctx context.Context, id uint) ([]*entity.TreeNode, error) {
	query := `
		SELECT 
			id, parent_id, user_id, file_type, extension, mime_type,
			upload_at, updated_at, name, description, size
		FROM tree_nodes
		WHERE user_id = $1
		ORDER BY file_type DESC, name
	`
	rows, err := r.db.Query(
		ctx,
		query,
		id,
	)
	if err != nil {
		return nil, fmt.Errorf("%w: find tree nodes by user id: %w", repository.ErrRepository, err)
	}
	defer rows.Close()

	nodes := make([]*entity.TreeNode, 0)

	for rows.Next() {
		var node entity.TreeNode

		err := rows.Scan(
			&node.ID,
			&node.ParentID,
			&node.UserID,
			&node.FileType,
			&node.Extension,
			&node.MimeType,
			&node.UploadAt,
			&node.UpdatedAt,
			&node.Name,
			&node.Description,
			&node.Size,
		)
		if err != nil {
			return nil, fmt.Errorf("%w: scan tree node: %w", repository.ErrRepository, err)
		}

		nodes = append(nodes, &node)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%w: iterate tree nodes: %w", repository.ErrRepository, err)
	}
	return nodes, nil
}

func (r *TreeNodeRepository) FindRoot(ctx context.Context, userID uint) (*entity.TreeNode, error) {
	query := `
		SELECT
			id, parent_id, user_id, file_type, extension, mime_type,
			upload_at, updated_at, name, description, size
		FROM tree_nodes
		WHERE user_id = $1 AND parent_id IS NULL
	`
	var node entity.TreeNode

	err := r.db.QueryRow(
		ctx,
		query,
		userID,
	).Scan(
		&node.ID,
		&node.ParentID,
		&node.UserID,
		&node.FileType,
		&node.Extension,
		&node.MimeType,
		&node.UploadAt,
		&node.UpdatedAt,
		&node.Name,
		&node.Description,
		&node.Size,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, repository.ErrTreeNodeNotFound
		}
		return nil, fmt.Errorf(
			"%w: find root: %w",
			repository.ErrRepository,
			err,
		)
	}
	return &node, nil
}

func (r *TreeNodeRepository) FindChildren(ctx context.Context, parentID *uuid.UUID, userID uint) ([]*entity.TreeNode, error) {
	var (
		query string
		rows  pgx.Rows
		err   error
	)
	if parentID == nil {
		query = `
			SELECT
				id, parent_id, user_id, file_type, extension, mime_type,
				upload_at, updated_at, name, description, size
			FROM tree_nodes
			WHERE parent_id IS NULL AND user_id = $1
			ORDER BY file_type DESC, name
		`

		rows, err = r.db.Query(ctx, query, userID)
	} else {
		query = `
			SELECT
				id, parent_id, user_id, file_type, extension, mime_type,
				upload_at, updated_at, name, description, size
			FROM tree_nodes
			WHERE parent_id = $1 AND user_id = $2
			ORDER BY file_type DESC, name
		`

		rows, err = r.db.Query(ctx, query, *parentID, userID)
	}
	if err != nil {
		return nil, fmt.Errorf("%w: find tree nodes by parent id: %w", repository.ErrRepository, err)
	}
	defer rows.Close()

	nodes := make([]*entity.TreeNode, 0)

	for rows.Next() {
		var node entity.TreeNode

		err := rows.Scan(
			&node.ID,
			&node.ParentID,
			&node.UserID,
			&node.FileType,
			&node.Extension,
			&node.MimeType,
			&node.UploadAt,
			&node.UpdatedAt,
			&node.Name,
			&node.Description,
			&node.Size,
		)
		if err != nil {
			return nil, fmt.Errorf("%w: scan tree node: %w", repository.ErrRepository, err)
		}

		nodes = append(nodes, &node)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%w: iterate tree nodes: %w", repository.ErrRepository, err)
	}
	return nodes, nil
}

func (r *TreeNodeRepository) UpdateName(ctx context.Context, id uuid.UUID, userID uint, name string) error {
	query := `
        UPDATE tree_nodes
        SET name = $3, updated_at = now()
        WHERE id = $1 AND user_id = $2;
    `
	result, err := r.db.Exec(
		ctx,
		query,
		id,
		userID,
		name,
	)
	if err != nil {
		return fmt.Errorf("%w: update name: %w",
			repository.ErrRepository,
			err,
		)
	}
	if result.RowsAffected() == 0 {
		return repository.ErrTreeNodeNotFound
	}
	return nil
}

func (r *TreeNodeRepository) UpdateDescription(ctx context.Context, id uuid.UUID, userID uint, description string) error {
	query := `
        UPDATE tree_nodes
        SET description = $3, updated_at = now()
        WHERE id = $1 AND user_id = $2
    `
	result, err := r.db.Exec(
		ctx,
		query,
		id,
		userID,
		description,
	)
	if err != nil {
		return fmt.Errorf("%w: update description: %w",
			repository.ErrRepository,
			err,
		)
	}
	if result.RowsAffected() == 0 {
		return repository.ErrTreeNodeNotFound
	}
	return nil
}

func (r *TreeNodeRepository) UpdateSize(ctx context.Context, id uuid.UUID, userID uint, size int64) error {
	query := `
        UPDATE tree_nodes
        SET size = $3, updated_at = now()
        WHERE id = $1 AND user_id = $2
    `
	result, err := r.db.Exec(
		ctx,
		query,
		id,
		userID,
		size,
	)
	if err != nil {
		return fmt.Errorf("%w: update size: %w",
			repository.ErrRepository,
			err,
		)
	}
	if result.RowsAffected() == 0 {
		return repository.ErrTreeNodeNotFound
	}
	return nil
}

func (r *TreeNodeRepository) UpdateFileMetadata(ctx context.Context, id uuid.UUID, userID uint, extension *string, mimeType *string, size int64) error {
	query := `
		UPDATE tree_nodes
		SET 
			extension = $3,
			mime_type = $4,
			size = $5,
			updated_at = now()
		WHERE id = $1 AND user_id = $2
	`
	result, err := r.db.Exec(
		ctx,
		query,
		id,
		userID,
		extension,
		mimeType,
		size,
	)
	if err != nil {
		return fmt.Errorf(
			"%w: update file metadata: %w",
			repository.ErrRepository,
			err,
		)
	}
	if result.RowsAffected() == 0 {
		return repository.ErrTreeNodeNotFound
	}
	return nil
}

func (r *TreeNodeRepository) Delete(ctx context.Context, id uuid.UUID, userID uint) error {
	query := `
		DELETE FROM tree_nodes
		WHERE id = $1 AND user_id = $2
	`
	result, err := r.db.Exec(
		ctx,
		query,
		id,
		userID,
	)
	if err != nil {
		return fmt.Errorf("%w: delete tree node: %w",
			repository.ErrRepository,
			err,
		)
	}
	if result.RowsAffected() == 0 {
		return repository.ErrTreeNodeNotFound
	}
	return nil
}

func (r *TreeNodeRepository) DeleteByUserID(ctx context.Context, userID uint) (int64, error) {
	query := `
		DELETE FROM tree_nodes
		WHERE user_id = $1
	`
	result, err := r.db.Exec(
		ctx,
		query,
		userID,
	)
	if err != nil {
		return 0, fmt.Errorf("%w: delete tree nodes by user id: %w",
			repository.ErrRepository,
			err,
		)
	}
	return result.RowsAffected(), nil
}

func (r *TreeNodeRepository) DeleteByParentID(ctx context.Context, parentID uuid.UUID, userID uint) (int64, error) {
	query := `
		DELETE FROM tree_nodes
		WHERE parent_id = $1 AND user_id = $2
	`
	result, err := r.db.Exec(
		ctx,
		query,
		parentID,
		userID,
	)
	if err != nil {
		return 0, fmt.Errorf("%w: delete tree nodes by parent id: %w",
			repository.ErrRepository,
			err,
		)
	}
	return result.RowsAffected(), nil
}
