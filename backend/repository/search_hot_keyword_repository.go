package repository

import (
	"container/heap"
	"dict/config"
	"dict/model"
	"log"
	"sync"
	"time"

	"gorm.io/gorm"
)

// 已有代码...

// SearchHotKeywordRepository 定义 SearchHotKeyword 存储库的接口
type SearchHotKeywordRepository interface {
	AddKeyword(keyword string)
	//FindSearchHotKeywordByID(id uint64) (model.SearchHotKeyword, error)
	//FindSearchHotKeywordByKeyword(keyword string) (model.SearchHotKeyword, error)
	//CreateSearchHotKeyword(keyword model.SearchHotKeyword) (model.SearchHotKeyword, error)
	//UpdateSearchHotKeyword(keyword model.SearchHotKeyword) (model.SearchHotKeyword, error)
	//DeleteSearchHotKeyword(id uint64) error
}

// NewSearchHotKeywordRepository 创建一个新的 SearchHotKeyword 存储库实例
func NewSearchHotKeywordRepository(size int) *searchHotKeywordRepository {
	return &searchHotKeywordRepository{config.GetDB(), &HotKeywordCache{
		entries: make(map[string]*KeywordEntry),
		heap:    make(KeywordHeap, 0),
		maxSize: size,
	}}
}

// searchHotKeywordRepository 实现 SearchHotKeywordRepository 接口
type searchHotKeywordRepository struct {
	db           *gorm.DB
	keywordCache *HotKeywordCache
}

// KeywordEntry 和 KeywordHeap 定义保持不变
type KeywordEntry struct {
	Keyword        string
	SearchCount    uint
	LastSearchedAt time.Time
	Index          int
}

type KeywordHeap []*KeywordEntry

func (h KeywordHeap) Len() int { return len(h) }
func (h KeywordHeap) Less(i, j int) bool {
	if h[i].SearchCount == h[j].SearchCount {
		return h[i].LastSearchedAt.Before(h[j].LastSearchedAt)
	}
	return h[i].SearchCount < h[j].SearchCount
}
func (h KeywordHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].Index = i
	h[j].Index = j
}
func (h *KeywordHeap) Push(x interface{}) {
	entry := x.(*KeywordEntry)
	entry.Index = len(*h)
	*h = append(*h, entry)
}
func (h *KeywordHeap) Pop() interface{} {
	old := *h
	n := len(old)
	entry := old[n-1]
	old[n-1] = nil
	entry.Index = -1
	*h = old[0 : n-1]
	return entry
}

// HotKeywordCache 定义
type HotKeywordCache struct {
	mu      sync.Mutex
	entries map[string]*KeywordEntry
	heap    KeywordHeap
	maxSize int
	db      *gorm.DB
}

func (c *HotKeywordCache) AddKeyword(keyword string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()
	if entry, exists := c.entries[keyword]; exists {
		entry.SearchCount++
		entry.LastSearchedAt = now
		heap.Fix(&c.heap, entry.Index)
	} else {
		entry = &KeywordEntry{
			Keyword:        keyword,
			SearchCount:    1,
			LastSearchedAt: now,
		}
		c.entries[keyword] = entry

		if len(c.heap) < c.maxSize {
			heap.Push(&c.heap, entry)
		} else if c.heap[0].SearchCount < entry.SearchCount ||
			(c.heap[0].SearchCount == entry.SearchCount && c.heap[0].LastSearchedAt.Before(now)) {
			old := heap.Pop(&c.heap).(*KeywordEntry)
			delete(c.entries, old.Keyword)
			heap.Push(&c.heap, entry)
		}
	}
}

func (c *HotKeywordCache) GetTopKeywords(n int) ([]KeywordEntry, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if len(c.heap) > 0 {
		if n > len(c.heap) {
			n = len(c.heap)
		}
		result := make([]KeywordEntry, n)
		tempHeap := make(KeywordHeap, len(c.heap))
		copy(tempHeap, c.heap)
		heap.Init(&tempHeap)

		for i := 0; i < n; i++ {
			if len(tempHeap) > 0 {
				result[i] = *heap.Pop(&tempHeap).(*KeywordEntry)
			}
		}
		return result, nil
	}

	var hotKeywords []model.SearchHotKeyword
	err := c.db.Order("search_count DESC, last_searched_at DESC").Find(&hotKeywords).Error
	if err != nil {
		return nil, err
	}

	result := make([]KeywordEntry, 0, n)
	for i, hk := range hotKeywords {
		if i >= n {
			break
		}
		result = append(result, KeywordEntry{
			Keyword:        hk.Keyword,
			SearchCount:    hk.SearchCount,
			LastSearchedAt: hk.LastSearchedAt,
		})
	}
	return result, nil
}

func (c *HotKeywordCache) SyncToDB() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	top10 := make(KeywordHeap, 0, 10)
	tempHeap := make(KeywordHeap, len(c.heap))
	copy(tempHeap, c.heap)
	heap.Init(&tempHeap)

	for i := 0; i < 10 && len(tempHeap) > 0; i++ {
		entry := heap.Pop(&tempHeap).(*KeywordEntry)
		top10 = append(top10, entry)
	}

	tx := c.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Exec("TRUNCATE TABLE search_hot_keywords").Error; err != nil {
		tx.Rollback()
		return err
	}

	for _, entry := range top10 {
		if err := tx.Create(&model.SearchHotKeyword{
			Keyword:        entry.Keyword,
			SearchCount:    entry.SearchCount,
			LastSearchedAt: entry.LastSearchedAt,
		}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (c *HotKeywordCache) StartSync() {
	ticker := time.NewTicker(1 * time.Minute)
	go func() {
		for range ticker.C {
			if err := c.SyncToDB(); err != nil {
				log.Printf("Failed to sync hot keywords to DB: %v", err)
			} else {
				log.Println("Hot keywords synced to DB (top 10)")
			}
		}
	}()
}

// FindSearchHotKeywordByID 根据 ID 查找 SearchHotKeyword
func (r *searchHotKeywordRepository) FindSearchHotKeywordByID(id uint64) (model.SearchHotKeyword, error) {
	var keyword model.SearchHotKeyword
	err := r.db.Where("id = ?", id).First(&keyword).Error
	if err != nil {
		return keyword, err
	}
	return keyword, nil
}

// FindSearchHotKeywordByKeyword 根据关键词查找 SearchHotKeyword
func (r *searchHotKeywordRepository) FindSearchHotKeywordByKeyword(keyword string) (model.SearchHotKeyword, error) {
	var hotKeyword model.SearchHotKeyword
	err := r.db.Where("keyword = ?", keyword).First(&hotKeyword).Error
	if err != nil {
		return hotKeyword, err
	}
	return hotKeyword, nil
}

// CreateSearchHotKeyword 创建一个新的 SearchHotKeyword 记录
func (r *searchHotKeywordRepository) CreateSearchHotKeyword(keyword model.SearchHotKeyword) (model.SearchHotKeyword, error) {
	err := r.db.Create(&keyword).Error
	if err != nil {
		return keyword, err
	}
	return keyword, nil
}

// UpdateSearchHotKeyword 更新一个已有的 SearchHotKeyword 记录
func (r *searchHotKeywordRepository) UpdateSearchHotKeyword(keyword model.SearchHotKeyword) (model.SearchHotKeyword, error) {
	err := r.db.Save(&keyword).Error
	if err != nil {
		return keyword, err
	}
	return keyword, nil
}

// DeleteSearchHotKeyword 根据 ID 删除一个 SearchHotKeyword 记录
func (r *searchHotKeywordRepository) DeleteSearchHotKeyword(id uint64) error {
	return r.db.Where("id = ?", id).Delete(&model.SearchHotKeyword{}).Error
}
