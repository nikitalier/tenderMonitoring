package service

import (
	"github.com/nikitalier/tenderMonitoring/pkg/models"
)

//UpdateFavorite ...
func (s *Service) UpdateFavorite(f models.Favorite) {
	if s.repository.FavoriteExists(f.TenderID, f.UserID) {
		// log.Println("1")
		s.repository.UpdateFavorite(f)
	} else {
		s.repository.AddNewFavorite(f)
		// log.Println("2")
	}
}

//GetFavoriteStatus ...
func (s *Service) GetFavoriteStatus(f models.Favorite) models.Favorite {
	var newF models.Favorite
	if s.repository.FavoriteExists(f.TenderID, f.UserID) {
		newF = s.repository.FindFavorite(f.TenderID, f.UserID)
		return newF
	}
	return newF
}
