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
	GetById(itemId int) error
	Delete(itemId int) error
	DeleteAll(userCode string) error
}

type ResizeOptions interface {
	Get(userCode string) (common.ResizeOptions, error)
	UpdateOrCreate(resizeOptions common.ResizeOptions, userCode string) error
	UpdateFinishTime(userCode string) error
	UpdateTotalAndCurrent(userCode string, totalCount, current int) error
	SaveError(userCode, errorMessage string) error
}

type Service struct {
	User
	File
	ResizeOptions
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User:          NewUserService(repos.User),
		File:          NewFileService(repos.File),
		ResizeOptions: NewResizeOptionsService(repos.ResizeOptions),
	}
}
