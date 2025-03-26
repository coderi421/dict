<script setup lang="ts">
import { useRoute } from "vue-router";

function toggleTheme(theme: string) {
  const root = document.documentElement;
  root.setAttribute("data-theme", theme);
  localStorage.setItem("theme", theme);
}
// 初始化时读取保存的主题
const savedTheme = localStorage.getItem("theme") || "system";
if (savedTheme !== "system") {
  document.documentElement.setAttribute("data-theme", savedTheme);
}

const route = useRoute();
console.log(route.query);
</script>

<template>
  <div>Home</div>
  <button @click="toggleTheme('light')">切换light主题</button>
  <button @click="toggleTheme('dark')">切换dark主题</button>
  <el-button type="primary">Primary</el-button>
  <el-button type="success">Success</el-button>
  <el-button type="info">Info</el-button>
</template>

<style scoped></style>
