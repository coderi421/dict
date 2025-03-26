package entity

type (
	SearchKeywordInput struct {
		Keyword string `json:"keyword" binding:"required"`
	}

	// CreateDictionaryInput 用于创建字典条目的输入结构体
	CreateDictionaryInput struct {
		Chinese            string `json:"chinese" binding:"required"`
		ChineseExplanation string `json:"chinese_explanation" binding:"required"`
		English            string `json:"english" binding:"required"`
		EnglishExplanation string `json:"english_explanation" binding:"required"`
		CategoryID         uint   `json:"category_id" binding:"required"`
		//Source             string `json:"source"`
		//Remark             string `json:"remark"`
	}

	// UpdateDictionaryInput 用于更新字典条目的输入结构体
	UpdateDictionaryInput struct {
		ID                 uint64 `json:"id" binding:"required"`
		Chinese            string `json:"chinese"`
		ChineseExplanation string `json:"chinese_explanation"`
		English            string `json:"english"`
		EnglishExplanation string `json:"english_explanation"`
		CategoryID         uint   `json:"category_id"`
		//Source             string `json:"source"`
		//Remark             string `json:"remark"`
	}

	// DeleteDictionaryInput 用于删除字典条目的输入结构体
	DeleteDictionaryInput struct {
		ID uint64 `json:"id" binding:"required"`
	}
)
