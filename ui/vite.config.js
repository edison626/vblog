import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
  server: {
    proxy:{
      // vite 类似与nginx， 后端 访问后端（没有浏览器的参与，咩有跨域的问题）
      '/api/vblog' : 'http://127.0.0.1:7080'
    }
  }
})
