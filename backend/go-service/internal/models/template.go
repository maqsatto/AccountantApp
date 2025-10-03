package models

import "time"

type Template struct {
	ID          uint    `gorm:"primaryKey;autoIncrement"`
	UserID      uint    `gorm:"not null"`
	Name        string  `gorm:"type:varchar(100);not null"`
	Description string  `gorm:"type:varchar(255)"`
	Amount      float64 `gorm:"type:numeric;not null"`
	CategoryID  uint
	AccountID   uint
	CreatedAt   time.Time
}
