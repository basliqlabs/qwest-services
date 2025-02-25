package postgresqlrefreshtoken

import (
	"context"
	"database/sql"
	"errors"
	"github.com/basliqlabs/qwest-services/internal/entity/tokenentity"
	"github.com/basliqlabs/qwest-services/internal/repository/postgresql"
	"github.com/basliqlabs/qwest-services/pkg/richerror"
)

type Repository struct {
	db *postgresql.DB
}

func New(db *postgresql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// Create stores a new refresh token
func (r *Repository) Create(ctx context.Context, token tokenentity.RefreshToken) (tokenentity.RefreshToken, error) {
	const op = "postgresqlrefreshtoken.Create"

	query := `
		INSERT INTO refresh_tokens (user_id, token, expires_at, created_at, revoked)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	stmt, err := r.db.Conn().PrepareContext(ctx, query)
	if err != nil {
		return tokenentity.RefreshToken{}, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected)
	}
	defer stmt.Close()

	var id int
	err = stmt.QueryRowContext(ctx,
		token.UserID,
		token.Token,
		token.ExpiresAt,
		token.CreatedAt,
		token.Revoked,
	).Scan(&id)

	if err != nil {
		return tokenentity.RefreshToken{}, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected)
	}

	token.ID = id
	return token, nil
}

// GetByToken retrieves a refresh token by its value
func (r *Repository) GetByToken(ctx context.Context, token string) (tokenentity.RefreshToken, error) {
	const op = "postgresqlrefreshtoken.GetByToken"

	query := `
		SELECT id, user_id, token, expires_at, created_at, revoked
		FROM refresh_tokens
		WHERE token = $1
	`

	stmt, err := r.db.Conn().PrepareContext(ctx, query)
	if err != nil {
		return tokenentity.RefreshToken{}, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected)
	}
	defer stmt.Close()

	var refreshToken tokenentity.RefreshToken
	err = stmt.QueryRowContext(ctx, token).Scan(
		&refreshToken.ID,
		&refreshToken.UserID,
		&refreshToken.Token,
		&refreshToken.ExpiresAt,
		&refreshToken.CreatedAt,
		&refreshToken.Revoked,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return tokenentity.RefreshToken{}, richerror.New(op).WithError(err).WithKind(richerror.KindNotFound)
		}
		return tokenentity.RefreshToken{}, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected)
	}

	return refreshToken, nil
}

// DeleteByUserID removes all refresh tokens for a user
func (r *Repository) DeleteByUserID(ctx context.Context, userID int) error {
	const op = "postgresqlrefreshtoken.DeleteByUserID"

	query := `
		DELETE FROM refresh_tokens
		WHERE user_id = $1
	`

	stmt, err := r.db.Conn().PrepareContext(ctx, query)
	if err != nil {
		return richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, userID)
	if err != nil {
		return richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected)
	}

	return nil
}

// DeleteByToken removes a specific refresh token
func (r *Repository) DeleteByToken(ctx context.Context, token string) error {
	const op = "postgresqlrefreshtoken.DeleteByToken"

	query := `
		DELETE FROM refresh_tokens
		WHERE token = $1
	`

	stmt, err := r.db.Conn().PrepareContext(ctx, query)
	if err != nil {
		return richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected)
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, token)
	if err != nil {
		return richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected)
	}

	if rowsAffected == 0 {
		return richerror.New(op).WithKind(richerror.KindNotFound).WithMessage("token not found")
	}

	return nil
}

// RevokeByToken marks a refresh token as revoked
func (r *Repository) RevokeByToken(ctx context.Context, token string) error {
	const op = "postgresqlrefreshtoken.RevokeByToken"

	query := `
		UPDATE refresh_tokens
		SET revoked = true
		WHERE token = $1
	`

	stmt, err := r.db.Conn().PrepareContext(ctx, query)
	if err != nil {
		return richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected)
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, token)
	if err != nil {
		return richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected)
	}

	if rowsAffected == 0 {
		return richerror.New(op).WithKind(richerror.KindNotFound).WithMessage("token not found")
	}

	return nil
}

func (r *Repository) GetByUserEmail(ctx context.Context, email string) (tokenentity.RefreshToken, error) {
	const op = "postgresqlrefreshtoken.GetByUserEmail"

	query := `
		SELECT id, user_id, token, expires_at, created_at, revoked
		FROM refresh_tokens
		WHERE user_id = (SELECT user_id FROM users WHERE email = $1)
		ORDER BY created_at DESC
		LIMIT 1
	`

	stmt, err := r.db.Conn().PrepareContext(ctx, query)
	if err != nil {
		return tokenentity.RefreshToken{}, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, email)
	if err != nil {
		return tokenentity.RefreshToken{}, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected)
	}
	defer rows.Close()

	var token tokenentity.RefreshToken
	rows.Next()

	err = rows.Scan(
		&token.ID,
		&token.UserID,
		&token.Token,
		&token.ExpiresAt,
		&token.CreatedAt,
		&token.Revoked,
	)

	if err != nil {
		return tokenentity.RefreshToken{}, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected)
	}

	if err := rows.Err(); err != nil {
		return tokenentity.RefreshToken{}, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected)
	}

	return token, nil
}
