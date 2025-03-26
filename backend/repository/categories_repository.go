package repository

import (
	"dict/config"
	"dict/model"
	"gorm.io/gorm"
)

// CategoryRepository 定义 Category 存储库接口
type CategoryRepository interface {
	FindCategoryByName(name string) (model.Category, error)
	// 可以添加更多方法
}

// categoryRepository 实现 CategoryRepository 接口
type categoryRepository struct {
	db *gorm.DB
}

// NewCategoryRepository 创建一个新的 Category 存储库实例
func NewCategoryRepository() *categoryRepository {
	return &categoryRepository{config.GetDB()}
}

// FindCategoryByName 根据名称查找 Category
func (r *categoryRepository) FindCategoryByName(name string) (model.Category, error) {
	var category model.Category
	err := r.db.Where("name = ?", name).First(&category).Error
	if err != nil {
		return category, err
	}
	return category, nil
}
