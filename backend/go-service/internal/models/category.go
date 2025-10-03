package models

import "time"

type Category struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	UserID    uint   `gorm:"not null"`
	Name      string `gorm:"type:varchar(100);not null"`
	Type      string `gorm:"type:varchar(50);not null"` // income / expense
	CreatedAt time.Time
}
