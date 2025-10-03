package models

import "time"

type User struct {
	ID           uint          `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string        `gorm:"type:varchar(100);not null" json:"name"`
	Email        string        `gorm:"uniqueIndex;type:varchar(100);not null" json:"email"`
	Password     string        `gorm:"type:varchar(255);not null" json:"password"`
	CreatedAt    time.Time     `json:"created_at"`
	Accounts     []Account     `gorm:"foreignKey:UserID" json:"accounts"`
	Transactions []Transaction `gorm:"foreignKey:UserID" json:"transactions"`
	Categories   []Category    `gorm:"foreignKey:UserID" json:"categories"`
	Templates    []Template    `gorm:"foreignKey:UserID" json:"templates"`
	Settings     []AppSettings `gorm:"foreignKey:UserID" json:"settings"`
}
