package models

import "time"

type Expense struct {
	Id         uint		`json:"id" gorm:"primaryKey"`
	Amount     float64 `json:"amount"`
	Category   string `json:"category"`
	Date       time.Time `json:"date"`
	Note       string	`json:"note"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
