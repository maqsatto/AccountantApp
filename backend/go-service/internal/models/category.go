package models

import "time"

type Category struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	Type      string    `gorm:"type:varchar(50);not null" json:"type"` // income or expense
	CreatedAt time.Time `json:"created_at"`
}
