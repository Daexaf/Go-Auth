// models/history.go

package models

import (
	"time"
)

type History struct {
	ID               string    `gorm:"primaryKey" json:"id"`
	TimeStamp        time.Time `json:"timestamp"`
	Amount           string    `json:"amount"`
	TransactionID    uint      `json:"transaction_id"`
	SenderAccountID  string    `json:"sender_account_id"`
	ReceiveAccountID string    `json:"receive_account_id"`
}
