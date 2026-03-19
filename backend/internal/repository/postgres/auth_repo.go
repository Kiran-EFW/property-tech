package postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/proptech/backend/internal/domain"
)

// AuthRepo implements service.AuthRepository using raw SQL against a pgx pool.
type AuthRepo struct {
	pool *pgxpool.Pool
}

// NewAuthRepo creates a new AuthRepo backed by the given connection pool.
func NewAuthRepo(pool *pgxpool.Pool) *AuthRepo {
	return &AuthRepo{pool: pool}
}

// CreateUser inserts a new user row and scans the returned columns back into
// the provided domain.User.
func (r *AuthRepo) CreateUser(ctx context.Context, user *domain.User) error {
	query := `
		INSERT INTO users (id, name, phone, email, password_hash, role, is_nri, avatar_url, created_at, updated_at)
		VALUES ($1, $2, $3, $4, '', $5, false, NULL, $6, $7)
		RETURNING id, phone, email, name, role, is_active, created_at, updated_at`

	row := r.pool.QueryRow(ctx, query,
		user.ID,
		user.Name,
		user.Phone,
		user.Email,
		user.Role,
		user.CreatedAt,
		user.UpdatedAt,
	)

	return row.Scan(
		&user.ID,
		&user.Phone,
		&user.Email,
		&user.Name,
		&user.Role,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
}

// GetUserByPhone retrieves a user by their phone number.
func (r *AuthRepo) GetUserByPhone(ctx context.Context, phone string) (*domain.User, error) {
	query := `
		SELECT id, phone, email, name, role, is_active, created_at, updated_at
		FROM users
		WHERE phone = $1`

	user := &domain.User{}
	row := r.pool.QueryRow(ctx, query, phone)

	err := row.Scan(
		&user.ID,
		&user.Phone,
		&user.Email,
		&user.Name,
		&user.Role,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("get user by phone: %w", err)
	}

	return user, nil
}

// GetUserByID retrieves a user by their unique identifier.
func (r *AuthRepo) GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	query := `
		SELECT id, phone, email, name, role, is_active, created_at, updated_at
		FROM users
		WHERE id = $1`

	user := &domain.User{}
	row := r.pool.QueryRow(ctx, query, id)

	err := row.Scan(
		&user.ID,
		&user.Phone,
		&user.Email,
		&user.Name,
		&user.Role,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("get user by id: %w", err)
	}

	return user, nil
}

// UpdateUser updates the name, email, and updated_at timestamp for the given user.
func (r *AuthRepo) UpdateUser(ctx context.Context, user *domain.User) error {
	query := `UPDATE users SET name = $1, email = $2, updated_at = $3 WHERE id = $4`

	_, err := r.pool.Exec(ctx, query,
		user.Name,
		user.Email,
		user.UpdatedAt,
		user.ID,
	)
	if err != nil {
		return fmt.Errorf("update user: %w", err)
	}

	return nil
}
