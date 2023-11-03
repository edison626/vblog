<template>
    <div>
      <!-- 顶部导航组件 -->
        <!-- 顶部 -->
        <TopBar></TopBar>
        <!-- 导航与内容 -->
          <div class="layout-content">
            <div class="menu">
              <a-menu
              :collapsed="collapsed"
              :style="{ width: '220px', height: '100%'}"
              :default-open-keys="['0', '1']"
              :default-selected-keys="['0_2']"
              show-collapse-button
              breakpoint="xl"
              @collapse="onCollapse"
              @menuItemClick="onMenuClick"
            >
              <a-sub-menu key="BackendBlogs">
                <template #icon><icon-apps></icon-apps></template>
                <template #title>博客管理</template>
                <a-menu-item key="BackendBlogs">文章列表</a-menu-item>
              </a-sub-menu>
              <a-sub-menu key="CommentList">
                <template #icon><icon-bug></icon-bug></template>
                <template #title>评论管理</template>
                <a-menu-item key="CommentList">评论列表</a-menu-item>
              </a-sub-menu>
            </a-menu>
            </div>
            <div class="content" :style="{width: contentWidth}">
              <router-view></router-view>
            </div>
          </div>
    </div>
  </template>
  

<script setup>
import TopBar from '../../components/TopBar.vue'
import { useRouter } from 'vue-router'
import { ref, computed } from 'vue'

//默认展开
const collapsed = ref(true)
const onCollapse = (v) => {
  collapsed.value = v
}
//根据状态计算内容区域的宽度
const contentWidth = computed(()=> {
  if (collapsed.value) {
    return 'calc(100vw - 48px)'
  } else {
    return "calc(100vw - 220px)"
  }
})


const router = useRouter();

const onMenuClick = (key) => {
  router.push({name: key})
}

</script>

<style lang="css" scoped>

.menu {
  border-right: 1px solid rgb(229, 230, 235);
}

.layout-content {
  display: flex;
  width: 100%;
  height: calc(100vh - 45px);
}
</style>