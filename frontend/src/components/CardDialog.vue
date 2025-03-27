<script setup lang="ts">
defineProps({
  modelValue: Boolean,
  currentResult: Object
})

defineEmits(['update:modelValue'])
</script>

<template>
  <div class="card-dialog">
    <!-- 新增对话框 -->
    <el-dialog
      :model-value="modelValue"
      @update:modelValue="$emit('update:modelValue', $event)"
      :title="'术语检索结果'"
      width="80%"
      style="margin-top: 12vh; "
    >
      <div class="dialog-container" style="max-height: 70vh; overflow-y: auto;">
        <div class="field">
          <h2>中文术语</h2>
          <p>{{ currentResult?.chinese }}</p>
        </div>

        <div class="field">
          <h2>中文释义</h2>
          <p>{{ currentResult?.chinese_explanation }}</p>
        </div>

        <!-- 以下字段需要确认数据源 -->
        <div class="field" v-if="currentResult?.english">
          <h2>英文术语</h2>
          <p>{{ currentResult.english }}</p>
        </div>
        <div class="field" v-if="currentResult?.english_explanation">
          <h2>英文释义</h2>
          <p>{{ currentResult.english_explanation }}</p>
        </div>

        <div class="field">
          <h2>所属领域</h2>
          <!--          <span class="domain-tag">-->
          <!--            {{ currentResult?.category_label || '未知分类' }}-->
          <!--          </span>-->
          <el-tag type="info" size="small" class="mt-4" style="display: table-cell;">
            所属领域：{{ currentResult?.category_label || '未知分类' }}
          </el-tag>
        </div>
      </div>
    </el-dialog>
  </div>
</template>
<style>
.dialog-container {
  margin: 0 auto;
  padding: 0 20px;
}

.card-dialog {
  .el-dialog .el-dialog__title {
    font-size: 30px;
    color: #2c3e50;
    font-weight: bold;
  }
}

.field {
  margin-bottom: 24px;
}

.field h2 {
  font-size: 18px;
  color: #34495e;
  margin-bottom: 8px;
  border-left: 4px solid #3498db;
  padding-left: 10px;
  text-align: left;
}

.field p {
  font-size: 16px;
  color: #2c3e50;
  text-align: justify;
  line-height: 1.6;
}

/* 新增滚动条美化 */
.dialog-container::-webkit-scrollbar {
  width: 8px;
}

.dialog-container::-webkit-scrollbar-thumb {
  background: #c1c1c1;
  border-radius: 4px;
}
</style>
