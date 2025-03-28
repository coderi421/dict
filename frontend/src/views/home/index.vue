<script setup lang="ts">
import { useRoute } from 'vue-router'
import { reactive, ref, onMounted } from 'vue'
import { Search } from '@element-plus/icons-vue'
import { getDictByKeyword, getAllCategory,getHotKeywords } from '@/api/home.js'
import SearchResult from '@/components/SearchResult.vue'
import HotSuggestions from '@/components/HotSuggestions.vue';
import { ElMessage } from 'element-plus'
import type { CategoryItem, HotKeywordItem } from '@/components/types.ts'
import type { ResultItem } from '@/views/home/types.ts'





// 定义 HotKeywordItem 类型


const route = useRoute()
console.log(route.query)

const input = ref('')
const select = ref('')
const categorys = ref<{ id: string; name: string }[]>([])
const searchButtonLoading = ref(false)
const results = reactive<ResultItem[]>([])
const suggestions = reactive<HotKeywordItem[]>([])

// 根据 category_id 获取分类标签
const getCategoryLabel = (categoryId: number) => {
  const category = categorys.value.find(item => Number(item.id) === categoryId)
  return category?.name || '未知分类'
}

// 搜索信息的方法
const searchInfoByKeyword = async (keyword: string, categoryId: string) => {
  try {
    searchButtonLoading.value = true
    const { data } = await getDictByKeyword({ keyword, category_id: categoryId })
    results.length = 0
    data.forEach((item: ResultItem) => {
      item.category_label = getCategoryLabel(item.category_id)
      results.push(item)
    })
  } catch (err: any) {
    ElMessage.error('Oops, 搜索数据失败, 请稍后再试')
  } finally {
    searchButtonLoading.value = false
  }
}

// 处理搜索按钮点击事件的方法
const handleSearch = () => {
  const keyword = input.value.trim()
  const categoryId = select.value

  searchInfoByKeyword(keyword, categoryId)
}

const handleSuggestionClick = (suggestion: string) => {
  input.value = suggestion
}
// 封装获取分类数据的方法
const fetchCategories = async () => {
  try {
    const response = await getAllCategory()
    const categoryItems: CategoryItem[] = response.data
    categorys.value = categoryItems.map((item: CategoryItem) => ({
      id: item.id.toString(),
      name: item.name
    }))
  } catch (error: any) {
    ElMessage.error('Oops, 获取分类数据失败, 请稍后再试')
  }
}
// 定义获取热门关键词的方法
const fetchHotKeywords = async () => {
  try {
    const response = await getHotKeywords();
    console.log('data', response)
    if (response?.length === 0) {
      return;
    }

    suggestions.length = 0;
    response.forEach((item: HotKeywordItem) => {
      suggestions.push(item);
    });
    console.log('suggestions', suggestions)
  } catch (err: any) {
    ElMessage.error('Oops, 获取热门关键词失败, 请稍后再试');
  }
};

onMounted(() => {
  // 获取热门搜索关键词
  fetchHotKeywords()
  // 设置定时器，每5分钟调用一次 fetchHotKeywords 方法
  setInterval(() => {
    fetchHotKeywords();
  }, 5 * 60 * 1000);

  // 获取分类数据
  fetchCategories()
  // 初始化，随机获取2条数据
  searchInfoByKeyword("", select.value)
})

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
            <el-option v-for="item in categorys" :key="item.id" :label="item.name" :value="item.id" />
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
      <HotSuggestions
        :suggestions="suggestions"
        @suggestion-click="handleSuggestionClick"
      />
      <SearchResult :results="results" :loading="searchButtonLoading" />
    </div>
  </div>
</template>

<style scoped>
.custom-col {
  padding: 0 5% !important;
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
  color: white !important;
  background-color: #16458a !important;
  text-align: center !important;
  font-size: large !important;
  border-top-left-radius: 0;
  border-bottom-left-radius: 0;
}

.search-input {
  height: 40px !important;
}

::v-deep .el-select__wrapper {
  height: 40px;
}
</style>