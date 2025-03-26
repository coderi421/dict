package model

import (
	"time"
)

// SearchHotKeyword 对应 search_hot_keywords 表的结构体
type SearchHotKeyword struct {
	ID             uint64    `gorm:"column:id;primaryKey;autoIncrement"`
	Keyword        string    `gorm:"column:keyword;type:varchar(255);not null;uniqueIndex:idx_keyword"`
	SearchCount    uint      `gorm:"column:search_count;not null;default:0"`
	LastSearchedAt time.Time `gorm:"column:last_searched_at;default:CURRENT_TIMESTAMP;autoUpdateTime"`
}
