package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	ID               int `gorm:"id"`
	SenderAccountID  string
	ReceiveAccountID string `gorm:"unique"`
	Amount           string
	TimeStamp        string
}
