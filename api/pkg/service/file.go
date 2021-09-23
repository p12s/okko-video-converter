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

func (s *FileService) Create(files []common.UploadedFile, userCode string) error {
	return s.repo.Create(files, userCode)
}

func (s *FileService) UpdateStatus(userCode, errorMess string, status common.ProcessStatus) error {
	return s.repo.UpdateStatus(userCode, errorMess, status)
}

func (s *FileService) GetByCode(userCode string) (common.File, error) {
	return s.repo.GetByCode(userCode)
}
