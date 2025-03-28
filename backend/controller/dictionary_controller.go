package controller

import (
	"dict/entity"
	"dict/helper"
	"dict/model"
	"dict/workflow"
	"github.com/gin-gonic/gin"
	"net/http"
)

// dictionaryController 字典控制器结构体
type dictionaryController struct {
	dictionaryService workflow.DictionaryService
	hotKeywordService workflow.SearchHotKeywordService
}

// NewDictionaryController 创建字典控制器实例
func NewDictionaryController(dictionaryService workflow.DictionaryService, hotKeywordController workflow.SearchHotKeywordService) *dictionaryController {
	return &dictionaryController{
		dictionaryService: dictionaryService,
		hotKeywordService: hotKeywordController,
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
func (d *dictionaryController) SearchDictionary(c *gin.Context) {
	keyword := c.Query("keyword")
	categoryId := c.Query("category_id")

	results, err := d.dictionaryService.SearchDictionary(keyword, categoryId)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		responseError := helper.APIResponse("服务器暂时不可用", http.StatusUnprocessableEntity, "fail", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, responseError)
		return
	}

	var formattedResults = []DictionaryFormatter{}

	var keywords = []string{}

	for _, result := range results {
		formattedResults = append(formattedResults, FormatDictionary(result))
		if keyword != "" {
			keywordType := helper.AnalyzeInputType(keyword)
			if keywordType == "pure_english" {
				keywords = append(keywords, result.Chinese)
				keywords = append(keywords, result.English)
			} else {
				keywords = append(keywords, result.Chinese)
			}
		}
	}
	go func() {
		if len(keywords) > 0 {
			//添加到搜索热词
			d.hotKeywordService.AddSearchHotKeyword(keywords)
		}
	}()
	////添加到搜索热词
	//d.hotKeywordService.AddSearchHotKeyword(keywords)

	response := helper.APIResponse("成功", http.StatusOK, "success", formattedResults)
	c.JSON(http.StatusOK, response)
}

// CreateDictionary 创建字典条目
func (d *dictionaryController) CreateDictionary(c *gin.Context) {
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

	createdDictionary, err := d.dictionaryService.CreateDictionary(dictionary)
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
func (d *dictionaryController) UpdateDictionary(c *gin.Context) {
	var input entity.UpdateDictionaryInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		responseError := helper.APIResponse("更新字典条目失败 #UPD001", http.StatusUnprocessableEntity, "fail", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, responseError)
		return
	}

	updatedDictionary, err := d.dictionaryService.UpdateDictionary(model.Dictionary{
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
func (d *dictionaryController) DeleteDictionary(c *gin.Context) {
	var input entity.DeleteDictionaryInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		responseError := helper.APIResponse("删除字典条目失败 #DEL001", http.StatusUnprocessableEntity, "fail", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, responseError)
		return
	}

	err = d.dictionaryService.DeleteDictionary(input.ID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		responseError := helper.APIResponse("删除字典条目失败 #DEL002", http.StatusBadRequest, "fail", errorMessage)
		c.JSON(http.StatusBadRequest, responseError)
		return
	}

	response := helper.APIResponse("字典条目删除成功", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
