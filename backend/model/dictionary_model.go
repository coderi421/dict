package model

import (
	"time"
)

// Dictionary 对应 dictionary 表的结构体
type Dictionary struct {
	ID                 uint64    `gorm:"column:id;primaryKey;autoIncrement"`
	Chinese            string    `gorm:"column:chinese;type:varchar(100);not null"`
	ChineseExplanation string    `gorm:"column:chinese_explanation;type:text;not null"`
	English            string    `gorm:"column:english;type:varchar(100);not null"`
	EnglishExplanation string    `gorm:"column:english_explanation;type:text;not null"`
	CategoryID         uint      `gorm:"column:category_id;not null"`
	Source             string    `gorm:"column:source;type:varchar(255)"`
	Remark             string    `gorm:"column:remark;type:text"`
	CreatedAt          time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
	UpdatedAt          time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;autoUpdateTime"`
	UpdatedBy          int       `gorm:"column:updated_by"`
}
