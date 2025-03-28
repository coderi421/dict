package controller

import (
	"dict/workflow"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SearchHotKeywordController 定义 SearchHotKeyword 控制器
type SearchHotKeywordController struct {
	service workflow.SearchHotKeywordService
}

// NewSearchHotKeywordController 创建一个新的 SearchHotKeyword 控制器实例
func NewSearchHotKeywordController(service workflow.SearchHotKeywordService) *SearchHotKeywordController {
	return &SearchHotKeywordController{
		service: service,
	}
}

// GetSearchHotKeyword 处理获取前 n 个热门搜索关键词的请求
func (s *SearchHotKeywordController) GetSearchHotKeyword(ctx *gin.Context) {
	//nStr := ctx.Query("n")
	//if nStr == "" {
	//	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing 'n' parameter"})
	//	return
	//}

	//n, err := strconv.Atoi(nStr)
	//if err != nil {
	//	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'n' parameter"})
	//	return
	//}
	keywords, err := s.service.GetSearchHotKeyword(8)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get search hot keywords"})
		return
	}

	ctx.JSON(http.StatusOK, keywords)
}

//
//// FindSearchHotKeywordByKeyword 处理根据关键词查找 SearchHotKeyword 的请求
//func (c *SearchHotKeywordController) FindSearchHotKeywordByKeyword(ctx *gin.Context) {
//	keyword := ctx.Param("keyword")
//	if keyword == "" {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing 'keyword' parameter"})
//		return
//	}
//
//	hotKeyword, err := c.service.FindSearchHotKeywordByKeyword(keyword)
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find search hot keyword"})
//		return
//	}
//
//	ctx.JSON(http.StatusOK, hotKeyword)
//}
//
//// UpdateSearchHotKeyword 处理更新已有的 SearchHotKeyword 记录的请求
//func (c *SearchHotKeywordController) UpdateSearchHotKeyword(ctx *gin.Context) {
//	var keyword model.SearchHotKeyword
//	if err := ctx.BindJSON(&keyword); err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
//		return
//	}
//
//	updatedKeyword, err := c.service.UpdateSearchHotKeyword(keyword)
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update search hot keyword"})
//		return
//	}
//
//	ctx.JSON(http.StatusOK, updatedKeyword)
//}
//
//// DeleteSearchHotKeyword 处理根据 ID 删除 SearchHotKeyword 记录的请求
//func (c *SearchHotKeywordController) DeleteSearchHotKeyword(ctx *gin.Context) {
//	idStr := ctx.Param("id")
//	if idStr == "" {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing 'id' parameter"})
//		return
//	}
//
//	id, err := strconv.ParseUint(idStr, 10, 64)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'id' parameter"})
//		return
//	}
//
//	err = c.service.DeleteSearchHotKeyword(id)
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete search hot keyword"})
//		return
//	}
//
//	ctx.JSON(http.StatusOK, gin.H{"message": "Search hot keyword deleted successfully"})
//}
