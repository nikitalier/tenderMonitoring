package service

import (
	"github.com/nikitalier/tenderMonitoring/pkg/models"
)

func (s *Service) GetSummary() (summary models.Summary) {
	summary.CountTenders = s.repository.CountAllTenders()
	summary.CountKeywords = s.repository.CountAllKeywords()
	summary.CountTendersByKeywords = s.repository.GetCountTendersByKeywords()
	summary.BestTenders = s.repository.GetBetsTenders()
	summary.ApprovedTenders = s.repository.GetApprovedTenders()

	return summary
}
