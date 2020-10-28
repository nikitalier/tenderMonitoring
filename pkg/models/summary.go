package models

type Summary struct {
	CountTenders           int
	CountKeywords          int
	CountTendersByKeywords []CountTendersByKeywords
	BestTenders            []Tender
	ApprovedTenders        []Tender
}

type CountTendersByKeywords struct {
	Keyword      string `db:"word"`
	CountTenders int    `db:"count"`
}
