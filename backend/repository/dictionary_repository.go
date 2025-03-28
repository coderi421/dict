package repository

import (
	"dict/config"
	"dict/helper"
	"dict/model"
	"gorm.io/gorm"
)

// DictionaryRepository 定义 Dictionary 存储库的接口
type DictionaryRepository interface {
	SearchDictionary(keyword string, categoryId string) ([]model.Dictionary, error)
	FindDictionaryByID(id uint64) (model.Dictionary, error)
	CreateDictionary(dictionary model.Dictionary) (model.Dictionary, error)
	UpdateDictionary(dictionary model.Dictionary) (model.Dictionary, error)
	DeleteDictionary(id uint64) error
}

// dictionaryRepository 实现 DictionaryRepository 接口
type dictionaryRepository struct {
	db *gorm.DB
}

// NewDictionaryRepository 创建一个新的 Dictionary 存储库实例
func NewDictionaryRepository() *dictionaryRepository {
	return &dictionaryRepository{config.GetDB()}
}

func (r *dictionaryRepository) SearchDictionary(keyword string, categoryId string) ([]model.Dictionary, error) {
	var results []model.Dictionary

	if keyword == "" {
		query := r.db.Table("dictionary")
		if categoryId != "" {
			query = query.Where("category_id = ?", categoryId)
		}
		err := query.Order("RAND()").Limit(2).Find(&results).Error
		if err != nil {
			return nil, err
		}
		return results, nil
	}

	inputType := helper.AnalyzeInputType(keyword)

	// 定义通用查询函数
	searchByField := func(field string) error {
		baseQuery := `
            SELECT chinese, chinese_explanation, english, english_explanation, category_id,
                   MATCH(` + field + `) AGAINST(? IN BOOLEAN MODE) AS ` + field + `_relevance,
                   MATCH(` + field + `_explanation) AGAINST(? IN BOOLEAN MODE) AS explanation_relevance
            FROM dictionary
            WHERE (MATCH(` + field + `) AGAINST(? IN BOOLEAN MODE)
               OR MATCH(` + field + `_explanation) AGAINST(? IN BOOLEAN MODE))
        `
		var args []interface{} = []interface{}{keyword, keyword, keyword, keyword}

		if categoryId != "" {
			baseQuery += " AND category_id = ?"
			args = append(args, categoryId)
		}

		baseQuery += " ORDER BY " + field + "_relevance DESC, explanation_relevance DESC"

		return r.db.Raw(baseQuery, args...).Scan(&results).Error
	}

	switch inputType {
	case "pure_chinese":
		if err := searchByField("chinese"); err != nil {
			return nil, err
		}
	case "pure_english":
		if err := searchByField("english"); err != nil {
			return nil, err
		}
	case "mixed":
		var chineseResults []model.Dictionary
		var englishResults []model.Dictionary

		if err := searchByField("chinese"); err != nil {
			return nil, err
		}
		chineseResults = append([]model.Dictionary{}, results...)

		results = nil // 清空 results 用于存储英文搜索结果
		if err := searchByField("english"); err != nil {
			return nil, err
		}
		englishResults = results

		// 合并结果
		results = append(results, chineseResults...)
		for _, result := range englishResults {
			// 避免重复添加
			exists := false
			for _, r := range results {
				if r.ID == result.ID {
					exists = true
					break
				}
			}
			if !exists {
				results = append(results, result)
			}
		}
	default:
		query := r.db.Table("dictionary").Where("chinese LIKE ? OR chinese_explanation LIKE ? OR english LIKE ?",
			"%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")
		if categoryId != "" {
			query = query.Where("category_id = ?", categoryId)
		}
		if err := query.Find(&results).Error; err != nil {
			return nil, err
		}
	}

	// 这部分可以用 goroutine 在controller层处理并发处理
	//if len(results) > 0 {
	//	fullKeyword := results[0].Chinese
	//	if inputType == "pure_english" || inputType == "mixed" {
	//		fullKeyword = results[0].English
	//	}
	//	cache.AddKeyword(fullKeyword)
	//}

	return results, nil

	//var results []model.Dictionary
	//inputType := AnalyzeInputType(keyword)

	//switch inputType {
	//case "pure_chinese":
	//	// 查询 chinese 和 chinese_explanation，标题匹配优先
	//	err := r.db.Raw(`
	//        SELECT chinese, chinese_explanation, english, english_explanation, category_id,
	//               MATCH(chinese) AGAINST(? IN BOOLEAN MODE) AS chinese_relevance,
	//               MATCH(chinese_explanation) AGAINST(? IN BOOLEAN MODE) AS explanation_relevance
	//        FROM dictionary
	//        WHERE MATCH(chinese) AGAINST(? IN BOOLEAN MODE)
	//           OR MATCH(chinese_explanation) AGAINST(? IN BOOLEAN MODE)
	//        ORDER BY chinese_relevance DESC, explanation_relevance DESC
	//    `, keyword, keyword, keyword, keyword).Scan(&results).Error
	//	if err != nil {
	//		return nil, err
	//	}
	//case "pure_english":
	//	// 查询 chinese 和 chinese_explanation，标题匹配优先
	//	err := r.db.Raw(`
	//        SELECT chinese, chinese_explanation, english, english_explanation, category_id,
	//               MATCH(english) AGAINST(? IN BOOLEAN MODE) AS english_relevance,
	//               MATCH(english_explanation) AGAINST(? IN BOOLEAN MODE) AS explanation_relevance
	//        FROM dictionary
	//        WHERE MATCH(english) AGAINST(? IN BOOLEAN MODE)
	//           OR MATCH(english_explanation) AGAINST(? IN BOOLEAN MODE)
	//        ORDER BY english_relevance DESC, explanation_relevance DESC
	//    `, keyword, keyword, keyword, keyword).Scan(&results).Error
	//	if err != nil {
	//		return nil, err
	//	}
	//case "pure_english":
	//	// 查询 english，匹配优先
	//	subQuery2 := db.Table("dictionary").
	//		Select("*, MATCH(english) AGAINST(? IN BOOLEAN MODE) AS english_relevance", keyword).
	//		Where("MATCH(english) AGAINST(? IN BOOLEAN MODE)", keyword)
	//
	//	err := subQuery2.Order("english_relevance DESC").Limit(10).Find(&results).Error
	//	if err != nil {
	//		return nil, err
	//	}
	//case "mixed":
	//	err := r.db.Where("MATCH(chinese) AGAINST(? IN BOOLEAN MODE) OR MATCH(english) AGAINST(? IN BOOLEAN MODE)", keyword, keyword).Find(&results).Error
	//	if err != nil {
	//		return nil, err
	//	}

	//default:
	//	err := r.db.Where("chinese LIKE ? OR chinese_explanation LIKE ? OR english LIKE ?",
	//		"%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%").Find(&results).Error
	//	if err != nil {
	//		return nil, err
	//	}
	//}

	//return results, nil
}

// FindDictionaryByID 根据 ID 查找 Dictionary
func (r *dictionaryRepository) FindDictionaryByID(id uint64) (model.Dictionary, error) {
	var dictionary model.Dictionary
	err := r.db.Where("id = ?", id).First(&dictionary).Error
	if err != nil {
		return dictionary, err
	}
	return dictionary, nil
}

// CreateDictionary 创建一个新的 Dictionary 记录
func (r *dictionaryRepository) CreateDictionary(dictionary model.Dictionary) (model.Dictionary, error) {
	err := r.db.Create(&dictionary).Error
	if err != nil {
		return dictionary, err
	}
	return dictionary, nil
}

// UpdateDictionary 更新一个已有的 Dictionary 记录
func (r *dictionaryRepository) UpdateDictionary(dictionary model.Dictionary) (model.Dictionary, error) {
	err := r.db.Save(&dictionary).Error
	if err != nil {
		return dictionary, err
	}
	return dictionary, nil
}

// DeleteDictionary 根据 ID 删除一个 Dictionary 记录
func (r *dictionaryRepository) DeleteDictionary(id uint64) error {
	return r.db.Where("id = ?", id).Delete(&model.Dictionary{}).Error
}
