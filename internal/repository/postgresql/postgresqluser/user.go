package postgresqluser

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	phonenumber "github.com/basliqlabs/qwest-services/pkg/mobile"

	"github.com/basliqlabs/qwest-services/internal/entity/userentity"
	"github.com/basliqlabs/qwest-services/pkg/errmsg"
	"github.com/basliqlabs/qwest-services/pkg/richerror"
)

func (d *DB) DoesIdentifierWithPasswordExist(ctx context.Context, identifier string, passwordHash string, identifierType string) (bool, error) {
	const op = "postgresqluser.DoesIdentifierWithPasswordExist"
	fmt.Println("query", passwordHash, identifierType)
	stmt, err := d.db.Conn().PrepareContext(ctx, fmt.Sprintf(`SELECT user_id FROM users WHERE %s=$1 AND password_hash=$2`, identifierType))
	if err != nil {
		return false, richerror.New(op).
			WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.CantScanQueryResult).
			WithError(err)
	}

	row := stmt.QueryRowContext(ctx, identifier, passwordHash)

	userId := new(int)
	err = row.Scan(userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, richerror.New(op).
			WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.CantScanQueryResult).
			WithError(err)
	}
	return true, nil
}

func (d *DB) scanUser(ctx context.Context, op string, query string, identifier string) (userentity.UserWithPasswordHash, bool, error) {
	stmt, err := d.db.Conn().PrepareContext(ctx, query)
	if err != nil {
		return userentity.UserWithPasswordHash{}, false, richerror.New(op).
			WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.CantScanQueryResult).
			WithError(err)
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, identifier)

	var user userentity.UserWithPasswordHash
	var mobile sql.NullString
	err = row.Scan(&user.UserName, &user.Email, &mobile, &user.PasswordHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return userentity.UserWithPasswordHash{}, false, nil
		}
		return userentity.UserWithPasswordHash{}, false, richerror.New(op).
			WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.CantScanQueryResult).
			WithError(err)
	}

	if mobile.Valid {
		user.Mobile = mobile.String
	}

	return user, true, nil
}

func (d *DB) FindUserByMobile(ctx context.Context, mobile string) (userentity.UserWithPasswordHash, bool, error) {
	const op = "postgresqluser.FindUserByMobile"

	normalizedMobile, err := phonenumber.NormalizePhoneNumber(mobile)
	if err != nil {
		return userentity.UserWithPasswordHash{}, false, err
	}

	user, found, err := d.scanUser(
		ctx,
		op,
		`SELECT username, email, mobile, password_hash FROM users WHERE mobile=$1`,
		normalizedMobile,
	)

	if err != nil {
		return userentity.UserWithPasswordHash{}, false, err
	}

	if !found {
		return userentity.UserWithPasswordHash{}, false, nil
	}

	// Ensure normalized mobile is used
	if user.Mobile != normalizedMobile {
		user.Mobile = normalizedMobile
	}

	return user, true, nil
}

func (d *DB) FindUserByEmail(ctx context.Context, email string) (userentity.UserWithPasswordHash, bool, error) {
	const op = "postgresqluser.FindUserByEmail"

	return d.scanUser(
		ctx,
		op,
		`SELECT username, email, mobile, password_hash FROM users WHERE email=$1`,
		email,
	)
}

func (d *DB) FindUserByUserName(ctx context.Context, username string) (userentity.UserWithPasswordHash, bool, error) {
	const op = "postgresqluser.FindUserByUserName"

	return d.scanUser(
		ctx,
		op,
		`SELECT username, email, mobile, password_hash FROM users WHERE username=$1`,
		username,
	)
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
