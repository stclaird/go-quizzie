package models

import (
	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	Text        string `json:"text"`
	Type        string `json:"type"`
	Category    string `json:"category"`
	Subcategory string `json:"subcategory"`
}

type Answer struct {
	gorm.Model
	Text       string `json:"text"`
	IsCorrect  bool   `json:"iscorrect"`
	QuestionID uint
	Question   Question `json:"question" gorm:"foreignKey:QuestionID"`
}

type ImportQuestion struct {
	Text        string `json:"text"`
	Type        string `json:"type"`
	Category    string `json:"category"`
	Subcategory string `json:"subcategory"`
	Answers     []struct {
		Text      string `json:"text"`
		IsCorrect bool   `json:"iscorrect"`
	} `json:"answers"`
}
