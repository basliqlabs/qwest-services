package postgresqluser

import (
	"context"
	"database/sql"
	"errors"
	phonenumber "github.com/basliqlabs/qwest-services/pkg/mobile"

	"github.com/basliqlabs/qwest-services/internal/entity/userentity"
	"github.com/basliqlabs/qwest-services/pkg/errmsg"
	"github.com/basliqlabs/qwest-services/pkg/richerror"
)

func (d *DB) DoesUserNameWithPasswordExist(ctx context.Context, username string, password string) (bool, error) {
	const op = "postgresqluser.DoesUserNameWithPasswordExist"
	stmt, err := d.db.Conn().PrepareContext(ctx, `SELECT user_id FROM users WHERE username=$1 AND password_hash=$2`)
	if err != nil {
		return false, richerror.New(op).
			WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.CantScanQueryResult).
			WithError(err)
	}

	row := stmt.QueryRowContext(ctx, username, password)

	userId := new(int)
	err = row.Scan(userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, richerror.New(op).
				WithKind(richerror.KindNotFound).
				WithMessage(errmsg.NotFound).
				WithError(err)
		}
		return false, richerror.New(op).
			WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.CantScanQueryResult).
			WithError(err)
	}
	return true, nil
}

func (d *DB) FindUserByMobile(ctx context.Context, mobile string) (userentity.UserWithPasswordHash, bool, error) {
	const op = "postgresqluser.FindUserByMobile"

	normalizedMobile, err := phonenumber.NormalizePhoneNumber(mobile)

	if err != nil {
		return userentity.UserWithPasswordHash{}, false, err
	}

	row := d.db.Conn().QueryRowContext(ctx,
		`SELECT username, email, password_hash FROM users WHERE mobile=$1`,
		normalizedMobile)

	var user userentity.UserWithPasswordHash
	err = row.Scan(&user.UserName, &user.Email, &user.PasswordHash)
	user.Mobile = normalizedMobile

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return userentity.UserWithPasswordHash{}, false, nil
		}
		return userentity.UserWithPasswordHash{}, false, richerror.New(op).
			WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.CantScanQueryResult).
			WithError(err)
	}
	return user, true, nil
}

func (d *DB) FindUserByEmail(ctx context.Context, email string) (userentity.UserWithPasswordHash, bool, error) {
	const op = "postgresqluser.FindUserByEmail"

	row := d.db.Conn().QueryRowContext(ctx,
		`SELECT username, email, password_hash FROM users WHERE email=$1`,
		email)

	var user userentity.UserWithPasswordHash
	err := row.Scan(&user.UserName, &user.Email, &user.PasswordHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return userentity.UserWithPasswordHash{}, false, nil
		}
		return userentity.UserWithPasswordHash{}, false, richerror.New(op).
			WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.CantScanQueryResult).
			WithError(err)
	}
	return user, true, nil
}

func (d *DB) FindUserByUserName(ctx context.Context, username string) (userentity.UserWithPasswordHash, bool, error) {
	const op = "postgresqluser.FindUserByUserName"
	row := d.db.Conn().QueryRowContext(ctx,
		`SELECT username, email, password_hash FROM users WHERE username=$1`,
		username)

	var user userentity.UserWithPasswordHash
	err := row.Scan(&user.UserName, &user.Email, &user.PasswordHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return userentity.UserWithPasswordHash{}, false, nil
		}
		return userentity.UserWithPasswordHash{}, false, richerror.New(op).
			WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.CantScanQueryResult).
			WithError(err)
	}
	return user, true, nil
}

func (d *DB) CreateUser(ctx context.Context, user userentity.UserWithPasswordHash) (int, error) {
	const op = "postgresqluser.CreateUser"

	stmt, err := d.db.Conn().PrepareContext(ctx, `INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3) RETURNING user_id`)

	if err != nil {
		return 0, richerror.New(op).
			WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.CantScanQueryResult).
			WithError(err)
	}

	row := stmt.QueryRowContext(ctx, user.UserName, user.Email, user.PasswordHash)

	userId := new(int)
	err = row.Scan(userId)

	if err != nil {
		return 0, richerror.New(op).
			WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.CantScanQueryResult).
			WithError(err)
	}
	return *userId, nil
}
