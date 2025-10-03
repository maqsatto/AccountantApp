package models

import "time"

type Account struct {
	ID           uint    `gorm:"primaryKey;autoIncrement"`
	UserID       uint    `gorm:"not null"`
	Name         string  `gorm:"type:varchar(100);not null"`
	Balance      float64 `gorm:"type:numeric;default:0"`
	Currency     string  `gorm:"type:varchar(10);default:'USD'"`
	CreatedAt    time.Time
	Transactions []Transaction `gorm:"foreignKey:AccountID"`
}
