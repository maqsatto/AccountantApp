package models

import "time"

type User struct {
	ID           uint   `gorm:"primaryKey;autoIncrement"`
	Name         string `gorm:"type:varchar(100);not null"`
	Email        string `gorm:"uniqueIndex;type:varchar(100);not null"`
	Password     string `gorm:"type:varchar(255);not null"`
	CreatedAt    time.Time
	Accounts     []Account     `gorm:"foreignKey:UserID"`
	Transactions []Transaction `gorm:"foreignKey:UserID"`
	Categories   []Category    `gorm:"foreignKey:UserID"`
	Templates    []Template    `gorm:"foreignKey:UserID"`
	Settings     AppSettings   `gorm:"foreignKey:UserID"`
}
