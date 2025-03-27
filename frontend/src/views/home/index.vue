<script setup lang="ts">
import { useRoute } from 'vue-router'
import { reactive, ref } from 'vue'
import { Search } from '@element-plus/icons-vue'
import { getDictByKeyword } from '@/api/home.js'
import SearchResult from '@/components/SearchResult.vue';

// 定义单个结果项的类型
interface ResultItem {
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

const route = useRoute()
console.log(route.query)

const input = ref('')
const select = ref('')
const categorys = ref([
  { value: '1', label: 'Restaurant' },
  { value: '2', label: 'Order No.' },
  { value: '3', label: 'Tel' }
])
const searchButtonLoading = ref(false)
const results = reactive<ResultItem[]>([])
const loading = ref(false)
const suggestions = reactive([
  '全面振兴',
  '乡村振兴',
  '装备制造',
  '文化遗产',
  '高质量发展'
])

// 根据 category_id 获取分类标签
const getCategoryLabel = (categoryId: number) => {
  const category = categorys.value.find(item => Number(item.value) === categoryId)
  return category?.label || '未知分类'
}

// 搜索信息的方法
const searchInfoByKeyword = async (keyword: string, categoryId: string) => {
  try {
    searchButtonLoading.value = true
    const { data } = await getDictByKeyword({ keyword, categoryId })
    results.length = 0
    data.forEach((item: ResultItem) => {
      item.category_label = getCategoryLabel(item.category_id)
      results.push(item)
    })
    console.log('results', results)
  } catch (err: any) {
    console.error('搜索出错:', err.message)
  } finally {
    searchButtonLoading.value = false
  }
}

// 处理搜索按钮点击事件的方法
const handleSearch = () => {
  const keyword = input.value.trim()
  const categoryId = select.value
  if (!keyword) {
    console.log('err')
    return
  }
  searchInfoByKeyword(keyword, categoryId)
}

const handleSuggestionClick = (suggestion:string) => {
  input.value = suggestion;
}
</script>

<template>
  <div>
    <div class="input-container">
      <el-input
        v-model="input"
        style="max-width: 700px; min-width: 300px; font-size: 18px;"
        placeholder="请输入检索关键字"
        class="input-with-select search-input"
        clearable
      >
        <template #prepend>
          <el-select v-model="select" clearable placeholder="分类" class="search-input" style="width: 140px;">
            <el-option v-for="item in categorys" :key="item.value" :label="item.label" :value="item.value" />
          </el-select>
        </template>
        <template #append>
          <el-button
            type="primary"
            class="search-button search-input"
            :loading="searchButtonLoading"
            :icon="Search"
            @click="handleSearch"
          />
        </template>
      </el-input>
      <!-- Search Section -->
      <el-row justify="center" class="mb-12">
        <el-col :span="18">
          <div style="margin-top: 20px">
            <span>热门推荐：</span>
            <el-button
              type="primary"
              plain
              round
              size="small"
              v-for="suggestion in suggestions"
              :key="suggestion"
              @click="handleSuggestionClick(suggestion)"
            >
              {{ suggestion }}
            </el-button>
          </div>
        </el-col>
      </el-row>
      <SearchResult :results="results" :loading="searchButtonLoading" />
    </div>
  </div>
</template>

<style scoped>
.custom-col {
  padding: 0 5%!important;
}
.custom-col .custom-card {
  text-align: left;
  margin-top: 1%;
  background-color: #fafafa;
}
.custom-col .custom-card .card-cell-h2 {
  font-size: 1.25rem;
  font-weight: 600;
  color: #1b3870;
}
.custom-col .custom-card .escription-text {
  display: -webkit-box;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 2;
  overflow: hidden;
  text-overflow: ellipsis;
}
.search-button {
  width: 60px;
  color: white!important;
  background-color: #16458a!important;
  text-align: center!important;
  font-size: large!important;
  border-top-left-radius: 0;
  border-bottom-left-radius: 0;
}
.search-input {
  height: 40px!important;
}
::v-deep .el-select__wrapper {
  height: 40px;
}
</style>