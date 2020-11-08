package service

import (
	"github.com/nikitalier/tenderMonitoring/pkg/repository"
	"github.com/rs/zerolog"
)

//Service -
type Service struct {
	repository *repository.Repository
	logger     *zerolog.Logger
}

//New - init service
func New(rep *repository.Repository, logger *zerolog.Logger) *Service {
	return &Service{
		repository: rep,
		logger:     logger,
	}
}
