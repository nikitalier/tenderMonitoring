package service

import (
	"github.com/nikitalier/tenderMonitoring/pkg/models"
	"github.com/nikitalier/tenderMonitoring/pkg/repository"
)

//Service -
type Service struct {
	repository *repository.Repository
}

//New - init service
func New(rep *repository.Repository) *Service {
	return &Service{
		repository: rep,
	}
}

//Test is just a test
func (s *Service) Test() {
	// log.Println(s.repository.FavExists(2, 5))
	// s.repository.FindFavorite(1, 1)
	var f models.Favorite
	f.TenderID = 1
	f.UserID = 5
	f.Status = false
	s.repository.UpdateFavorite(f)
}
