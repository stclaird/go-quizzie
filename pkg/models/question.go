package models

type Question struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Question string `json:"question"`
}
