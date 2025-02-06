package postgresqluser

import (
	"context"
	"database/sql"

	"github.com/basliqlabs/qwest-services/entity/userentity"
	"github.com/basliqlabs/qwest-services/pkg/errmsg"
	"github.com/basliqlabs/qwest-services/pkg/richerror"
)

func (d *DB) DoesUserNameWithPasswordExist(ctx context.Context, username string, password string) (bool, error) {
	const op = "postgresqluser.DoesUserNameWithPasswordExist"
	row := d.db.Conn().QueryRowContext(ctx,
		`SELECT user_id FROM users WHERE username=? AND password_hash=?`,
		username, password)
	userId := new(int)
	err := row.Scan(userId)
	if err != nil {
		if err == sql.ErrNoRows {
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
	return userentity.UserWithPasswordHash{}, false, nil
}
func (d *DB) FindUserByEmail(ctx context.Context, email string) (userentity.UserWithPasswordHash, bool, error) {
	return userentity.UserWithPasswordHash{}, false, nil
}
func (d *DB) FindUserByUserName(ctx context.Context, username string) (userentity.UserWithPasswordHash, bool, error) {
	return userentity.UserWithPasswordHash{}, false, nil
}
