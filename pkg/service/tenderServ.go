package service

import (
	"github.com/nikitalier/tenderMonitoring/pkg/models"
)

//GetAllTenders ...
func (s *Service) GetAllTenders() (tenders []models.Tender) {
	return s.repository.GetAllTenders()
}

//GetTender ...
func (s *Service) GetTender(id int) (tender models.Tender) {
	return s.repository.GetTenderByID(id)
}

//GetTenderStatus ...
func (s *Service) GetTenderStatus(id int) (status models.TenderStatus) {
	status = s.repository.GetTenderStatusByID(id)
	// log.Println(status)
	return status
}

func (s *Service) CreateTenderStatus(id int) {
	if !s.repository.StatusExists(id) {
		s.repository.CreateTenderStatus(id)
	}
}

func (s *Service) UpdateTenderStatus(status models.TenderStatus) {
	s.repository.UpdateTenderStatus(status)
}
