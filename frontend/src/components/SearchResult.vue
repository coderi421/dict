<script setup lang="ts">
import { ref, defineProps } from 'vue'
import CardDialog from './CardDialog.vue'
import { ResultItem } from '@/path/to/Home';

const props = defineProps<{
  results: ResultItem[];
  loading: boolean;
}>()
const dialogVisible = ref(false)
const currentResult = ref()

const handleShowDetail = (item: ResultItem) => {
  currentResult.value = item
  dialogVisible.value = true
}
</script>
<template>
  <div v-if="loading" class="loading-container">
    <el-skeleton  animated  style="margin-top: 40px;">
      <template #template>
        <el-skeleton-item variant="text" style="height: 240px; width: 90%;text-align: center" />
      </template>
    </el-skeleton>
  </div>
  <div v-else>
    <!-- 添加对话框组件 -->
    <CardDialog
      v-model="dialogVisible"
      :current-result="currentResult"
    />
    <!-- Results Section -->
    <el-row v-if="results.length" :gutter="20" style="margin-top: 20px">
      <el-col
        :span="24"
        v-for="result in results"
        :key="result.id"
        class="custom-col"
      >
        <el-card shadow="hover" class="custom-card" @click="handleShowDetail(result)">
          <h2 class="card-cell-h2">{{ result.chinese }}</h2>
          <p class="escription-text" @click.stop>{{ result.chinese_explanation }}</p>
          <el-tag type="info" size="small" class="mt-4" style="display: table-cell;" @click.stop>
            所属领域：{{ result.category_label || '未知分类' }}
          </el-tag>
        </el-card>
      </el-col>
    </el-row>
    <el-empty v-else description="未检测到结果" style="width: 100%;" />
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


</style>