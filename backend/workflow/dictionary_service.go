package workflow

import (
	"dict/model"
	"dict/repository"
	"errors"
)

// DictionaryService 定义 Dictionary 服务的接口
type DictionaryService interface {
	SearchDictionary(keyword string) ([]model.Dictionary, error)
	FindDictionaryByID(id uint64) (model.Dictionary, error)
	CreateDictionary(dictionary model.Dictionary) (model.Dictionary, error)
	UpdateDictionary(dictionary model.Dictionary) (model.Dictionary, error)
	DeleteDictionary(id uint64) error
}

// dictionaryService 实现 DictionaryService 接口
type dictionaryService struct {
	service repository.DictionaryRepository
}

// NewDictionaryService 创建一个新的 Dictionary 服务实例
func NewDictionaryService(repository repository.DictionaryRepository) *dictionaryService {
	return &dictionaryService{
		service: repository,
	}
}

// SearchDictionary 搜索字典条目
func (s *dictionaryService) SearchDictionary(keyword string) ([]model.Dictionary, error) {
	results, err := s.service.SearchDictionary(keyword)
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, errors.New("未找到相关字典条目")
	}
	return results, nil
}

// FindDictionaryByID 根据 ID 查找字典条目
func (s *dictionaryService) FindDictionaryByID(id uint64) (model.Dictionary, error) {
	dictionary, err := s.service.FindDictionaryByID(id)
	if err != nil {
		return model.Dictionary{}, err
	}
	if dictionary.ID == 0 {
		return model.Dictionary{}, errors.New("未找到指定 ID 的字典条目")
	}
	return dictionary, nil
}

// CreateDictionary 创建一个新的字典条目
func (s *dictionaryService) CreateDictionary(dictionary model.Dictionary) (model.Dictionary, error) {
	return s.service.CreateDictionary(dictionary)
}

// UpdateDictionary 更新一个已有的字典条目
func (s *dictionaryService) UpdateDictionary(dictionary model.Dictionary) (model.Dictionary, error) {
	return s.service.UpdateDictionary(dictionary)
}

// DeleteDictionary 根据 ID 删除一个字典条目
func (s *dictionaryService) DeleteDictionary(id uint64) error {
	return s.service.DeleteDictionary(id)
}
