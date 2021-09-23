package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/p12s/okko-video-converter/api/common"
)

// User - интерфейс для работы с пользователем
type User interface {
	CreateUser(userCode uuid.UUID) error
	GetUser(code uuid.UUID) (common.User, error)
}

// File - интерфейс для работы с файлом
type File interface {
	GetAll(userCode string) ([]common.File, error)
	DeleteAll(userCode string) error
	Create(files []common.UploadedFile, userCode string) error
	UpdateStatus(userCode, errorMess string, status common.ProcessStatus) error
	GetByCode(userCode string) (common.File, error)
}

// Repository - репозиторий
type Repository struct {
	User
	File
}

// NewRepository - конструктор
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: NewUserPostgres(db),
		File: NewFilePostgres(db),
	}
}
