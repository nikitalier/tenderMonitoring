package service

import (
	"log"
	"time"

	"github.com/nikitalier/tenderMonitoring/pkg/models"
)

const (
	dateFormat = "02.01.2006"
)

//GetAllKeyWords ...
func (s *Service) GetAllKeyWords() (keywords []models.Keyword) {
	keywords = s.repository.GetAllKeyWords()
	for i := 0; i < len(keywords); i++ {
		keywords[i].AddDateString = keywords[i].AddDateTime.Format(dateFormat)
	}
	return keywords
}

//DeleteKeywords ...
func (s *Service) DeleteKeywords(ids string) bool {
	return s.repository.DeleteKeyWordsByIDs(ids)
}

//AddKeyword ...
// func (s *Service) AddKeyword(keyword models.Keyword) bool {
// 	return s.repository.AddKeyword(keyword)
// }

//AddKeyword ...
func (s *Service) AddKeyword(keyword models.Keyword) (id int, result bool) {
	var err error
	keyword.AddDateTime, err = time.Parse(dateFormat, keyword.AddDateString)
	if err != nil {
		log.Println(err)
	}
	id, result = s.repository.AddKeyword(keyword)
	return id, result
}
