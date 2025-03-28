<script setup lang="ts">
import { defineProps, defineEmits } from 'vue'
import type { HotKeywordItem } from '@/components/types.ts'

defineProps({
  suggestions: {
    type: Array as () => HotKeywordItem[],
    default: () => []
  }
})

const emits = defineEmits(['suggestion-click'])
</script>
<template>
  <el-row justify="center" v-if="suggestions?.length > 0" class="mb-12">
    <el-col :span="18">
      <div style="margin-top: 20px">
        <span>热门推荐：</span>
        <el-tooltip
          :content="suggestion.Keyword"
          placement="bottom"
          effect="light"
          :hide-after="0"
          v-for="suggestion in suggestions"
          :key="suggestion.Index"
        >
          <el-button
            v-if="suggestion.Keyword"
            type="primary"
            plain
            round
            @click="$emit('suggestion-click', suggestion.Keyword)"
            size="small"
            style="min-width: 15px;"
          >
            {{ suggestion?.Keyword?.length > 8 ? suggestion?.Keyword?.slice(0, 8) + '...' : suggestion.Keyword }}
          </el-button>
        </el-tooltip>
      </div>
    </el-col>
  </el-row>
</template>

