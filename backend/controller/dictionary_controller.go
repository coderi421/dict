package controller

import (
	"dict/entity"
	"dict/helper"
	"dict/model"
	"dict/workflow"
	"net/http"

	"github.com/gin-gonic/gin"
)

// dictionaryController 字典控制器结构体
type dictionaryController struct {
	dictionaryService workflow.DictionaryService
}

// NewDictionaryController 创建字典控制器实例
func NewDictionaryController(dictionaryService workflow.DictionaryService) *dictionaryController {
	return &dictionaryController{
		dictionaryService: dictionaryService,
	}
}

// DictionaryFormatter 字典格式化结构体
type DictionaryFormatter struct {
	ID                 uint64 `json:"id"`
	Chinese            string `json:"chinese"`
	ChineseExplanation string `json:"chinese_explanation"`
	English            string `json:"english"`
	EnglishExplanation string `json:"english_explanation"`
	CategoryID         uint   `json:"category_id"`
	Source             string `json:"source"`
	Remark             string `json:"remark"`
}

// FormatDictionary 格式化字典数据
func FormatDictionary(dictionary model.Dictionary) DictionaryFormatter {
	return DictionaryFormatter{
		ID:                 dictionary.ID,
		Chinese:            dictionary.Chinese,
		ChineseExplanation: dictionary.ChineseExplanation,
		English:            dictionary.English,
		EnglishExplanation: dictionary.EnglishExplanation,
		CategoryID:         dictionary.CategoryID,
		Source:             dictionary.Source,
		Remark:             dictionary.Remark,
	}
}

// SearchDictionary 搜索字典条目
func (h *dictionaryController) SearchDictionary(c *gin.Context) {
	keyword := c.Query("keyword")
	if keyword == "" {
		errorMessage := gin.H{"errors": "搜索关键词不能为空"}
		responseError := helper.APIResponse("搜索失败 #SRCH001", http.StatusUnprocessableEntity, "fail", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, responseError)
		return
	}

	results, err := h.dictionaryService.SearchDictionary(keyword)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		responseError := helper.APIResponse("搜索失败 #SRCH002", http.StatusUnprocessableEntity, "fail", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, responseError)
		return
	}

	var formattedResults = []DictionaryFormatter{}
	for _, result := range results {
		formattedResults = append(formattedResults, FormatDictionary(result))
	}

	response := helper.APIResponse("搜索成功", http.StatusOK, "success", formattedResults)
	c.JSON(http.StatusOK, response)
}

// CreateDictionary 创建字典条目
func (h *dictionaryController) CreateDictionary(c *gin.Context) {
	var input entity.CreateDictionaryInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		responseError := helper.APIResponse("创建字典条目失败 #CRT001", http.StatusUnprocessableEntity, "fail", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, responseError)
		return
	}

	dictionary := model.Dictionary{
		Chinese:            input.Chinese,
		ChineseExplanation: input.ChineseExplanation,
		English:            input.English,
		EnglishExplanation: input.EnglishExplanation,
		CategoryID:         input.CategoryID,
	}

	createdDictionary, err := h.dictionaryService.CreateDictionary(dictionary)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		responseError := helper.APIResponse("创建字典条目失败 #CRT002", http.StatusBadRequest, "fail", errorMessage)
		c.JSON(http.StatusBadRequest, responseError)
		return
	}

	response := helper.APIResponse("字典条目创建成功", http.StatusOK, "success", FormatDictionary(createdDictionary))
	c.JSON(http.StatusOK, response)
}

// UpdateDictionary 更新字典条目
func (h *dictionaryController) UpdateDictionary(c *gin.Context) {
	var input entity.UpdateDictionaryInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		responseError := helper.APIResponse("更新字典条目失败 #UPD001", http.StatusUnprocessableEntity, "fail", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, responseError)
		return
	}

	updatedDictionary, err := h.dictionaryService.UpdateDictionary(model.Dictionary{
		ID:                 input.ID,
		Chinese:            input.Chinese,
		ChineseExplanation: input.ChineseExplanation,
		English:            input.English,
		EnglishExplanation: input.EnglishExplanation,
		CategoryID:         input.CategoryID,
	})
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		responseError := helper.APIResponse("更新字典条目失败 #UPD002", http.StatusBadRequest, "fail", errorMessage)
		c.JSON(http.StatusBadRequest, responseError)
		return
	}

	response := helper.APIResponse("字典条目更新成功", http.StatusOK, "success", FormatDictionary(updatedDictionary))
	c.JSON(http.StatusOK, response)
}

// DeleteDictionary 删除字典条目
func (h *dictionaryController) DeleteDictionary(c *gin.Context) {
	var input entity.DeleteDictionaryInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		responseError := helper.APIResponse("删除字典条目失败 #DEL001", http.StatusUnprocessableEntity, "fail", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, responseError)
		return
	}

	err = h.dictionaryService.DeleteDictionary(input.ID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		responseError := helper.APIResponse("删除字典条目失败 #DEL002", http.StatusBadRequest, "fail", errorMessage)
		c.JSON(http.StatusBadRequest, responseError)
		return
	}

	response := helper.APIResponse("字典条目删除成功", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
