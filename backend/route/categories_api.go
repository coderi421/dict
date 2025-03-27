package route

import (
	"dict/controller"
	"dict/workflow"

	"github.com/gin-gonic/gin"
)

func RouteCategory(route *gin.Engine, userService workflow.UserService, categoryService workflow.CategoryService) {
	//authService := config.NewServiceAuth()
	categoryController := controller.NewCategoryController(categoryService)
	//categoryMiddleware := middleware.AuthMiddlewareUser(authService, userService)

	api := route.Group("/api/v1/")
	api.GET("category/all", categoryController.GetAllCategories)
	// 你可以根据实际需求添加更多的 category 相关 API，如创建、更新、删除等
	// api.POST("category/create", categoryController.CreateCategory)
	// api.POST("category/update", categoryMiddleware, categoryController.UpdateCategory)
	// api.POST("category/delete", categoryMiddleware, categoryController.DeleteCategory)
}
