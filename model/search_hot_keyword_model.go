package model

import "time"

type SearchHotKeyword struct {
	ID             uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Keyword        string    `gorm:"type:varchar(100);unique;not null" json:"keyword"`
	SearchCount    uint      `gorm:"type:int unsigned;not null;default:0" json:"search_count"`
	LastSearchedAt time.Time `gorm:"autoUpdateTime" json:"last_searched_at"`
}
