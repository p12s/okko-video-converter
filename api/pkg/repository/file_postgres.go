package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/p12s/okko-video-converter/api/common"
)

type FilePostgres struct {
	db *sqlx.DB
}

func NewFilePostgres(db *sqlx.DB) *FilePostgres {
	return &FilePostgres{db: db}
}

func (r *FilePostgres) GetAll(userCode string) ([]common.File, error) {
	var items []common.File
	query := fmt.Sprintf(`SELECT f.path, f.name, f.user_id, f.kilo_byte_size, f.prev_image, f.status, f.error_message 
									FROM %s f INNER JOIN %s u on u.id = f.user_id
									WHERE u.code = $1`,
		fileTable, usersTable)
	if err := r.db.Select(&items, query, userCode); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *FilePostgres) DeleteAll(userCode string) error {
	query := fmt.Sprintf(`DELETE FROM %s f USING %s u WHERE u.id = f.user_id AND u.code = $1`,
		fileTable, usersTable)
	_, err := r.db.Exec(query, userCode)
	return err
}

func (r *FilePostgres) Create(files []common.UploadedFile, userCode string) error { // TODO в будущем можно добавить и сохранять ошибку загрузки, если она есть
	if len(files) == 0 {
		return nil
	}

	query := fmt.Sprintf("INSERT INTO %s (path, name, user_id, kilo_byte_size, prev_image, status) VALUES ", fileTable)

	for i, item := range files {
		query += fmt.Sprintf("('%s', '%s', (SELECT id FROM users WHERE code='%s'), %d, '%s', %d)",
			item.Path, item.Name, userCode, item.KiloByteSize, item.PrevImage, common.STARTED)
		if i < len(files)-1 {
			query += ","
		}
	}

	_, err := r.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (r *FilePostgres) UpdateStatus(userCode, errorMess string, status common.ProcessStatus) error {
	var query string
	if errorMess != "" {
		query = fmt.Sprintf(`UPDATE %s AS f SET status=$1, error_message=$2 FROM %s AS u WHERE f.user_id = u.id AND u.code = $3`,
			fileTable, usersTable)
	} else {
		query = fmt.Sprintf(`UPDATE %s AS f SET status=$1, error_message=$2 FROM %s AS u WHERE f.user_id = u.id AND u.code = $3`,
			fileTable, usersTable)
	}

	_, err := r.db.Exec(query, status, errorMess, userCode)

	return err
}

func (r *FilePostgres) GetByCode(userCode string) (common.File, error) {
	var item common.File
	query := fmt.Sprintf(`SELECT f.path, f.name, f.user_id, f.kilo_byte_size, f.prev_image, f.status 
									FROM %s f INNER JOIN %s u on u.id = f.user_id
									WHERE u.code = $1`,
		fileTable, usersTable)
	if err := r.db.Get(&item, query, userCode); err != nil {
		return item, err
	}

	return item, nil
}
