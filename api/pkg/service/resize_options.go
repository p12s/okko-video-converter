package service

import (
	"github.com/p12s/okko-video-converter/api/common"
	"github.com/p12s/okko-video-converter/api/pkg/repository"
)

type ResizeOptionsService struct {
	repo repository.ResizeOptions
}

func NewResizeOptionsService(repo repository.ResizeOptions) *ResizeOptionsService {
	return &ResizeOptionsService{repo: repo}
}

func (s *ResizeOptionsService) Get(userCode string) (common.ResizeOptions, error) {
	return s.repo.Get(userCode)
}

func (s *ResizeOptionsService) UpdateOrCreate(resizeOptions common.ResizeOptions, userCode string) error {
	return s.repo.UpdateOrCreate(resizeOptions, userCode)
}

func (s *ResizeOptionsService) UpdateFinishTime(userCode string) error {
	return s.repo.UpdateFinishTime(userCode)
}

func (s *ResizeOptionsService) UpdateTotalAndCurrent(userCode string, totalCount, current int) error {
	return s.repo.UpdateTotalAndCurrent(userCode, totalCount, current)
}

func (s *ResizeOptionsService) SaveError(userCode, errorMessage string) error {
	return s.repo.SaveError(userCode, errorMessage)
}
