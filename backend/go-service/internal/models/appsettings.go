package models

type AppSettings struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	UserID uint   `json:"user_id"`
	Key    string `gorm:"type:varchar(100)" json:"key"`
	Value  string `gorm:"type:varchar(255)" json:"value"`
}
