package models

import "time"

type Transaction struct {
	ID          uint    `gorm:"primaryKey;autoIncrement"`
	UserID      uint    `gorm:"not null"`
	AccountID   uint    `gorm:"not null"`
	Amount      float64 `gorm:"type:numeric;not null"`
	Description string  `gorm:"type:varchar(255)"`
	CategoryID  uint
	Date        time.Time `gorm:"not null"`
	CreatedAt   time.Time
}
