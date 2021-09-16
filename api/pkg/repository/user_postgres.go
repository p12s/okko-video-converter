package repository

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/p12s/okko-video-converter/api/common"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) CreateUser(userCode uuid.UUID) error {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (code) values ($1) RETURNING id", usersTable)
	row := r.db.QueryRow(query, userCode)
	if err := row.Scan(&id); err != nil {
		return err
	}

	return nil
}

func (r *UserPostgres) GetUser(code uuid.UUID) (common.User, error) {
	var user common.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE code=$1", usersTable)
	err := r.db.Get(&user, query, code.String())

	return user, err
}
