package model

import "time"

// Dictionary
type Dictionary struct {
	ID                 uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Chinese            string    `gorm:"type:varchar(100);not null" json:"chinese"`
	ChineseExplanation string    `gorm:"type:text;not null" json:"chinese_explanation"`
	English            string    `gorm:"type:varchar(100);not null" json:"english"`
	EnglishExplanation string    `gorm:"type:text;not null" json:"english_explanation"`
	CategoryId         uint      `gorm:"type:int;not null" json:"category_id"`
	Source             string    `gorm:"type:varchar(255)" json:"source"`
	Remark             string    `gorm:"type:text" json:"remark"`
	CreatedAt          time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt          time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
