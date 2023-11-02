<template>
    <div class="top-bar">
      <div style="margin-left: 12px">老喻的博客系统</div>
      <div>
        <a-button type="text"  @click="login"  v-if="!state.is_login">登录</a-button>
        <a-button type="text" @click="logout" v-if="state.is_login">
          <span style="margin-right: 12px;">退出</span>
          <icon-export />
        </a-button>
      </div>
    </div>
  </template>
  
  <script setup>
  import { state } from '../stores/app'
  import { useRouter } from 'vue-router'

  //console.log(state.value);
  //console.log(state.value.is_login);
  const router = useRouter()
  const logout = () => {
    state.value.is_login = false
    //调用后端的Token销毁接口（作业）

    
  }
  const login = () => {
    router.push({name: 'Login'})
  }

  </script>
  
  <style lang="css" scoped>
  .top-bar {
    height: 45px;
    border-bottom: 1px solid rgb(229, 230, 235);
    display: flex;
    align-items: center;
    justify-content: space-between;
  }
  
  /* 调整第三方组件库的样式, 使用css :deep函数 */
  :deep(.arco-btn) {
    height: 100%;
  }
  </style>
  