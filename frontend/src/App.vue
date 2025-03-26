<script setup lang="ts">
import Header from './views/header/index.vue'
import Footer from './views/footer/index.vue'
import { useRouter } from 'vue-router'

const router = useRouter()
var aa = '1111'
console.log(aa)
</script>

<template>
  <!--  <div>-->
  <!--    <a href="https://vite.dev" target="_blank">-->
  <!--      <img src="/vite.svg" class="logo" alt="Vite logo" />-->
  <!--    </a>-->
  <!--    <a href="https://vuejs.org/" target="_blank">-->
  <!--      <img src="./assets/vue.svg" class="logo vue" alt="Vue logo" />-->
  <!--    </a>-->
  <!--  </div>-->
  <!--  <HelloWorld msg="Vite + Vue" />-->
  <!--  <button @click="() => router.push({ path: '/home' })">首页</button>-->
  <!--  <button @click="() => router.push({ path: '/mine' })">个人中心</button>-->
  <!--  <router-link to="/home">首页2</router-link>-->


<!--  <div class="common-layout">-->
<!--    <el-container>-->
<!--      <el-header class="header">-->
<!--        <Header />-->
<!--      </el-header>-->
<!--      <el-main>-->
<!--        <router-view />-->
<!--      </el-main>-->
<!--      <el-footer>Footer</el-footer>-->
<!--    </el-container>-->
<!--  </div>-->
  <div class="main">
    <el-container class="main-content">
<!--      <el-aside :width="isCollapse ? '60px' : '210px'">-->
<!--        <nav-menu :collapse="isCollapse" />-->
<!--      </el-aside>-->
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
      <el-footer class="footer">
        <Footer />
      </el-footer>
    </el-container>
  </div>
</template>


<style scoped>
.header {
  background: linear-gradient(to right, #2c3e50, #4a6278);
  color: white;
  padding: 10px 20px;
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
  position: fixed;
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
