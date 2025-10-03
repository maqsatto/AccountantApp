package models

import "time"

type AppSettings struct {
	ID              uint   `gorm:"primaryKey;autoIncrement"`
	UserID          uint   `gorm:"uniqueIndex;not null"`
	DefaultCurrency string `gorm:"type:varchar(10);default:'USD'"`
	Theme           string `gorm:"type:varchar(20);default:'light'"`
	Notifications   bool   `gorm:"default:true"`
	CreatedAt       time.Time
}
