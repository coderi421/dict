export type HotKeywordItem = {
  Keyword: string;
  SearchCount: number;
  LastSearchedAt: string;
  Index: number;
};

export // 定义 CategoryItem 类型
type CategoryItem = {
  id: number;
  name: string;
};