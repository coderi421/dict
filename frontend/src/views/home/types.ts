// 定义单个结果项的类型
export type ResultItem = {
  id: number
  chinese: string
  chinese_explanation: string
  english: string
  english_explanation: string
  category_id: number
  source: string
  remark: string
  category_label?: string // 新增可选属性，用于存储分类标签
}