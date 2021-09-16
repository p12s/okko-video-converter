package service

import (
	"github.com/p12s/okko-video-converter/api/common"
	"github.com/p12s/okko-video-converter/api/pkg/repository"
)

type FileService struct {
	repo repository.File
}

func NewFileService(repo repository.File) *FileService {
	return &FileService{repo: repo}
}

func (s *FileService) GetAll(userCode string) ([]common.File, error) {
	return s.repo.GetAll(userCode)
}

func (s *FileService) DeleteAll(userCode string) error {
	return s.repo.DeleteAll(userCode)
}

// Create - создание
func (s *FileService) Create(files []common.UploadedFile, userCode string) error {
	return s.repo.Create(files, userCode)
}

// GetById - получение по id
func (s *FileService) GetById(itemId int) error {
	return s.repo.GetById(itemId)
}

// Delete - удаление
func (s *FileService) Delete(itemId int) error {
	return s.repo.Delete(itemId)
}
