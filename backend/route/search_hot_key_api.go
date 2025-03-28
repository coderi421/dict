package route

import (
	"dict/controller"
	"dict/workflow"

	"github.com/gin-gonic/gin"
)

// RouteSearchHotKey 定义搜索热门关键词相关的 API 路由
func RouteSearchHotKey(route *gin.Engine, hotKeyService workflow.SearchHotKeywordService) {
	hotKeyController := controller.NewSearchHotKeywordController(hotKeyService)

	api := route.Group("/api/v1/")

	// 获取热门关键词
	api.GET("hotkeys", hotKeyController.GetSearchHotKeyword)

}
