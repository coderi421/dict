package controller

import (
	"dict/helper"
	"dict/model"
	"dict/workflow"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CategoryController 定义 Category 控制器的结构体
type CategoryController struct {
	service workflow.CategoryService
}

// NewCategoryController 创建一个新的 Category 控制器实例
func NewCategoryController(service workflow.CategoryService) *CategoryController {
	return &CategoryController{service: service}
}

// CategoryFormatter 分类格式化结构体
type CategoryFormatter struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	// 可以根据实际 model.Category 结构体添加更多字段
}

// FormatCategory 格式化分类数据
func FormatCategory(category model.Category) CategoryFormatter {
	return CategoryFormatter{
		ID:   category.ID,
		Name: category.Name,
		// 可以根据实际 model.Category 结构体添加更多字段映射
	}
}

// GetAllCategories 获取所有的 category
func (c *CategoryController) GetAllCategories(ctx *gin.Context) {
	categories, err := c.service.GetAllCategories()
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		responseError := helper.APIResponse("搜索失败 #SRCH002", http.StatusUnprocessableEntity, "fail", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, responseError)
		return
	}

	var formattedResults = []CategoryFormatter{}
	for _, result := range categories {
		formattedResults = append(formattedResults, FormatCategory(result))
	}

	response := helper.APIResponse("搜索成功", http.StatusOK, "success", formattedResults)
	ctx.JSON(http.StatusOK, response)
}
