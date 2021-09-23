package service

import (
	"github.com/google/uuid"
	"github.com/p12s/okko-video-converter/api/common"
	"github.com/p12s/okko-video-converter/api/pkg/repository"
)

type User interface {
	CreateUser() (uuid.UUID, error)
	GenerateToken(userCode uuid.UUID) (string, error)
	ParseToken(accessToken string) (uuid.UUID, error)
}

type File interface {
	GetAll(userCode string) ([]common.File, error)
	Create(files []common.UploadedFile, userCode string) error
	UpdateStatus(userCode, errorMess string, status common.ProcessStatus) error
	DeleteAll(userCode string) error
	GetByCode(userCode string) (common.File, error)
}

type Service struct {
	User
	File
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repos.User),
		File: NewFileService(repos.File),
	}
}
