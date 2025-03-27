<script setup lang="ts">
import Header from './views/header/index.vue'
import Footer from './views/footer/index.vue'
import { useRouter } from 'vue-router'

const router = useRouter()
var aa = '1111'
console.log(aa)
</script>

<template>
  <div class="app-wrapper">
    <div class="main">
      <el-container class="main-content">
        <el-container>
          <el-header class="header">
            <Header />
          </el-header>
          <el-main class="page-content">
            <div class="page-info">
              <router-view></router-view>
            </div>
          </el-main>
        </el-container>
      </el-container>
    </div>
    <el-footer class="footer">
      <Footer />
    </el-footer>
  </div>
</template>


<style scoped>

.app-wrapper {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
}

.main {
  flex: 1; /* Make the main content area take up the remaining space */
}
.footer {
  background: #f7f7f7;
  color: #888;
  text-align: center;
  padding: 20px 0;
  box-shadow: inset 0 2px 10px rgba(0, 0, 0, 0.05);
  height: 60px;
  flex-shrink: 0;
}

.header {
  /* 使用导入的背景图片 */
  background-image: url('@/assets/background4.jpg');
  background-size: cover;
  background-position: center;
  /* 添加半透明的黑色遮罩层 */
  background-color: rgba(44, 62, 80, 0.3); /* #2c3e50 的 RGBA 表示，透明度 0.7 */
  background-blend-mode: overlay;

  color: white;
  text-align: center;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  height: 180px;
  flex-shrink: 0;
}

.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}

.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}

.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}

.main {
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
}



.main-content,
.page {
  height: 100%;
}

.page-content {
  height: calc(100%);

  .page-info {
    background-color: #ffffff;
    border-radius: 5px;
  }
}

.footer {
  background: #f7f7f7;
  color: #888;
  text-align: center;
  padding: 20px 0;
  box-shadow: inset 0 2px 10px rgba(0, 0, 0, 0.05);
  height: 60px;
  flex-shrink: 0;
}

</style>
