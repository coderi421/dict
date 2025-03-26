package model

import (
	"time"
)

// User 对应 users 表的结构体
type User struct {
	ID        uint      `gorm:"column:id;primaryKey;autoIncrement"`
	Username  string    `gorm:"column:username;type:varchar(255);not null;unique"`
	Password  string    `gorm:"column:password;type:varchar(255);not null"`
	Role      int       `gorm:"column:role;not null"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;autoUpdateTime"`
	UpdatedBy int       `gorm:"column:updated_by"`
}
