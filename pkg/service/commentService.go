package service

import (
	"github.com/nikitalier/tenderMonitoring/pkg/models"
)

//GetAllComments ...
func (s *Service) GetAllComments(tenderID int) (comments []models.Comment) {
	comments = s.repository.GetAllCommentsByTenderID(tenderID)
	return comments
}

//AddNewComment ...
func (s *Service) AddNewComment(comment models.Comment) {
	s.repository.AddNewComment(comment)
}
