package workflow

import (
	"dict/model"
	"dict/repository"
)

// CategoryService 定义 Category 服务的接口
type CategoryService interface {
	GetAllCategories() ([]model.Category, error)
}

// categoryService 实现 CategoryService 接口
type categoryService struct {
	repo repository.CategoryRepository
}

// NewCategoryService 创建一个新的 Category 服务实例
func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{repo: repo}
}

// GetAllCategories 获取所有的 category
func (s *categoryService) GetAllCategories() ([]model.Category, error) {
	return s.repo.GetAllCategories()
}
