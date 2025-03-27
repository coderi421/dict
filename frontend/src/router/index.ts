import home from "@/views/home/index.vue";
import { createRouter, createWebHistory } from "vue-router";

const routes = [
  { path: "/", name: "home", component: home },
  // { path: "/mine", name: "mine", component: mine },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;