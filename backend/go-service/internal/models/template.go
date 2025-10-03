package models

import "time"

type Template struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID      uint      `gorm:"not null" json:"user_id"`
	Name        string    `gorm:"type:varchar(100);not null" json:"name"`
	Description string    `gorm:"type:varchar(255)" json:"description"`
	Amount      float64   `gorm:"type:numeric;not null" json:"amount"`
	CategoryID  uint      `json:"category_id"`
	AccountID   uint      `json:"account_id"`
	CreatedAt   time.Time `json:"created_at"`
}
