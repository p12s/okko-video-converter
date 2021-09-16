package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/p12s/okko-video-converter/api/common"
)

type ResizeOptionsPostgres struct {
	db *sqlx.DB
}

func NewResizeOptionsPostgres(db *sqlx.DB) *ResizeOptionsPostgres {
	return &ResizeOptionsPostgres{db: db}
}

func (r *ResizeOptionsPostgres) Get(userCode string) (common.ResizeOptions, error) {
	var item common.ResizeOptions
	query := fmt.Sprintf(`SELECT r.id, r.user_id, r.options, r.start_date, 
									r.status, r.total_count, r.current
									FROM %s r INNER JOIN %s u on u.id = r.user_id
									WHERE u.code = $1`,
		resizeOptionsTable, usersTable)
	err := r.db.Get(&item, query, userCode)

	return item, err
}

func (r *ResizeOptionsPostgres) UpdateOrCreate(resizeOptions common.ResizeOptions, userCode string) error {
	var user common.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE code=$1", usersTable)

	err := r.db.Get(&user, query, userCode)
	if err != nil {
		return err
	}
	if user.Id == 0 {
		return errors.New("User with that code is not exists")
	}

	var exists bool
	query = fmt.Sprintf("SELECT exists (SELECT id FROM %s WHERE user_id=$1)", resizeOptionsTable)
	err = r.db.QueryRow(query, user.Id).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if exists {
		var existsResizeOptionsId int
		query = fmt.Sprintf("SELECT id FROM %s WHERE user_id=$1", resizeOptionsTable)
		err = r.db.Get(&existsResizeOptionsId, query, user.Id)
		if err != nil {
			return err
		}
		query = fmt.Sprintf("UPDATE %s SET options=$1, start_date=NOW(), finish_date=NULL, status=$2 WHERE id=$3",
			resizeOptionsTable)
		_, err := r.db.Exec(query, resizeOptions.Options, resizeOptions.Status, existsResizeOptionsId)

		if err != nil {
			return err
		}

	} else {
		query = fmt.Sprintf("INSERT INTO %s (user_id, options, status) VALUES ($1, $2, $3)", resizeOptionsTable)
		_, err := r.db.Exec(query, user.Id, resizeOptions.Options, int(resizeOptions.Status))

		if err != nil {
			return err
		}
	}
	return nil
}

func (r *ResizeOptionsPostgres) UpdateFinishTime(userCode string) error {
	query := fmt.Sprintf(`UPDATE %s AS r SET finish_date = now(), status = $1 FROM %s AS u 
		WHERE r.user_id = u.id AND u.code = $2`,
		resizeOptionsTable, usersTable)
	_, err := r.db.Exec(query, common.FINISHED, userCode)
	return err
}

func (r *ResizeOptionsPostgres) UpdateTotalAndCurrent(userCode string, totalCount, current int) error {
	query := fmt.Sprintf(`UPDATE %s AS r SET total_count=$1, current=$2 FROM %s AS u
		WHERE r.user_id = u.id AND u.code = $3`,
		resizeOptionsTable, usersTable)
	_, err := r.db.Exec(query, totalCount, current, userCode)
	return err
}

func (r *ResizeOptionsPostgres) SaveError(userCode, errorMessage string) error {
	query := fmt.Sprintf(`UPDATE %s AS r SET finish_date = null, status = $1, error_message = $2 FROM %s AS u 
		WHERE r.user_id = u.id AND u.code = $3`,
		resizeOptionsTable, usersTable)
	_, err := r.db.Exec(query, common.ERROR, errorMessage, userCode)
	return err
}
