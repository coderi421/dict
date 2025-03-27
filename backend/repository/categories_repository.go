package repository

import (
	"dict/config"
	"dict/model"
	"gorm.io/gorm"
)

// CategoryRepository 定义 Category 存储库的接口
type CategoryRepository interface {
	GetAllCategories() ([]model.Category, error)
}

// categoryRepository 实现 CategoryRepository 接口
type categoryRepository struct {
	db *gorm.DB
}

// NewCategoryRepository 创建一个新的 Category 存储库实例
func NewCategoryRepository() CategoryRepository {
	return &categoryRepository{db: config.GetDB()}
}

// GetAllCategories 获取所有的 category
func (r *categoryRepository) GetAllCategories() ([]model.Category, error) {
	var categories []model.Category
	err := r.db.Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}
