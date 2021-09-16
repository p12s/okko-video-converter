package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/p12s/okko-video-converter/api/common"
)

type User interface {
	CreateUser(userCode uuid.UUID) error
	GetUser(code uuid.UUID) (common.User, error)
}

type File interface {
	GetAll(userCode string) ([]common.File, error)
	DeleteAll(userCode string) error
	Create(files []common.UploadedFile, userCode string) error
	GetById(itemId int) error
	Delete(itemId int) error
}

type ResizeOptions interface {
	Get(userCode string) (common.ResizeOptions, error)
	UpdateOrCreate(resizeOptions common.ResizeOptions, userCode string) error
	UpdateFinishTime(userCode string) error
	UpdateTotalAndCurrent(userCode string, totalCount, current int) error
	SaveError(userCode, errorMessage string) error
}

type Repository struct {
	User
	File
	ResizeOptions
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User:          NewUserPostgres(db),
		File:          NewFilePostgres(db),
		ResizeOptions: NewResizeOptionsPostgres(db),
	}
}
