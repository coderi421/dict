package workflow

import (
	"context"
	"dict/helper"
	"dict/repository"
	"sync"
	"time"
)

// SearchHotKeywordService 定义 SearchHotKeyword 服务的接口
type SearchHotKeywordService interface {
	GetSearchHotKeyword(n int) ([]repository.KeywordEntry, error)
	AddSearchHotKeyword(keywords []string)
	//FindSearchHotKeywordByKeyword(keyword string) (model.SearchHotKeyword, error)
	//UpdateSearchHotKeyword(keyword model.SearchHotKeyword) (model.SearchHotKeyword, error)
	//DeleteSearchHotKeyword(id uint64) error
}

// searchHotKeywordService 实现 SearchHotKeywordService 接口
type searchHotKeywordService struct {
	repository repository.SearchHotKeywordRepository
}

// NewSearchHotKeywordService 创建一个新的 SearchHotKeyword 服务实例
func NewSearchHotKeywordService(repo repository.SearchHotKeywordRepository) *searchHotKeywordService {
	return &searchHotKeywordService{
		repository: repo,
	}
}

// GetSearchHotKeyword 获取前 n 个热门搜索关键词
func (s *searchHotKeywordService) GetSearchHotKeyword(n int) ([]repository.KeywordEntry, error) {
	return s.repository.GetSearchHotKeyword(n)
}

// AddSearchHotKeyword 添加一个新的 SearchHotKeyword 记录
func (s *searchHotKeywordService) AddSearchHotKeyword(keywords []string) {
	if len(keywords) == 0 {
		return
	}
	var wg sync.WaitGroup
	wg.Add(1)

	// 创建一个带超时的 context，防止 goroutine 长时间运行
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 启动一个 goroutine 调用 AddSearchHotKeyword 方法
	go func(ctx context.Context, keywords []string) {
		defer wg.Done()
		select {
		case <-ctx.Done():
			// 处理超时或者上下文取消的情况
			helper.Log.Error(ctx, "AddSearchHotKeyword goroutine cancelled: %v", ctx.Err())
		default:
			s.repository.AddSearchHotKeyword(keywords)
		}
	}(ctx, keywords)

	wg.Wait() // 等待 goroutine 完成
}

//// FindSearchHotKeywordByKeyword 根据关键词查找 SearchHotKeyword
//func (s *searchHotKeywordService) FindSearchHotKeywordByKeyword(keyword string) (model.SearchHotKeyword, error) {
//	return s.repository.FindSearchHotKeywordByKeyword(keyword)
//}
//
//// UpdateSearchHotKeyword 更新一个已有的 SearchHotKeyword 记录
//func (s *searchHotKeywordService) UpdateSearchHotKeyword(keyword model.SearchHotKeyword) (model.SearchHotKeyword, error) {
//	return s.repository.UpdateSearchHotKeyword(keyword)
//}
//
//// DeleteSearchHotKeyword 根据 ID 删除一个 SearchHotKeyword 记录
//func (s *searchHotKeywordService) DeleteSearchHotKeyword(id uint64) error {
//	return s.repository.DeleteSearchHotKeyword(id)
//}
