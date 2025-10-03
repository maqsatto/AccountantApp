package models

import "time"

type Transaction struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID      uint      `gorm:"not null" json:"user_id"`
	AccountID   uint      `gorm:"not null" json:"account_id"`
	Amount      float64   `gorm:"type:numeric;not null" json:"amount"`
	Description string    `gorm:"type:varchar(255)" json:"description"`
	CategoryID  uint      `json:"category_id"`
	Date        time.Time `gorm:"not null" json:"date"`
	CreatedAt   time.Time `json:"created_at"`
}
