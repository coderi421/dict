import vue from "@vitejs/plugin-vue";
import { resolve } from "path";
import { defineConfig } from 'vite'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import { ElementPlusResolver } from 'unplugin-vue-components/resolvers'


const pathResolve = (path: string): string => resolve(process.cwd(), path);

// https://vite.dev/config/
export default defineConfig({
  resolve: {
    alias: {
      // '@': path.resolve(__dirname, 'src')
      "@": pathResolve("src"),
    },
  },
  plugins: [vue(),
    AutoImport({
      resolvers: [ElementPlusResolver()],
    }),
    Components({
      resolvers: [ElementPlusResolver()],
    }),
  ],
});


