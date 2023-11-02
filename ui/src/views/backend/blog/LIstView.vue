<template>
    <div class="page">
      <!-- 页头 -->
      <a-breadcrumb>
        <a-breadcrumb-item>博客管理</a-breadcrumb-item>
        <a-breadcrumb-item>文章管理</a-breadcrumb-item>
      </a-breadcrumb>
      <!-- 内容区 -->
      <!-- 表格操作区 -->
      <div class="table-op">
        <div>
          <a-button type="primary" size="small">创建文章</a-button>
        </div>
        <div>
          <a-input :style="{width:'320px'}" placeholder="Please enter something" allow-clear />
        </div>
      </div>
      <!-- 表格内容 -->
      <div>
        <!-- 使用后端分页,  自己适配 -->
        <a-table :data="blogs.items" :pagination="false">
          <template #columns>
            <a-table-column title="编号" data-index="id"></a-table-column>
            <a-table-column title="标题" data-index="title"></a-table-column>
            <a-table-column title="作者" data-index="author"></a-table-column>
            <a-table-column title="状态" data-index="status"></a-table-column>
            <!-- 使用dayjs来处理时间 -->
            <a-table-column title="状态">
              <template #cell="{ record }">
                  {{ dayjs.unix(record.created_at).format('YYYY-MM-DD HH:mm')  }}
                </template>
            </a-table-column>
          </template>
        </a-table>
        <!-- 适配后端分页 -->
        <div style="margin-top: 6px;">
          <a-pagination 
            :total="blogs.total" 
            show-total 
            show-jumper 
            :page-size-options="[2, 10, 20, 30, 50]"
            show-page-size
            @page-size-change="onPageSizeChange"
            @change="onPageNumberChange"
            />
        </div>
  
      </div>
    </div>
  </template>
  
  
<script setup>
    import { onBeforeMount, ref } from 'vue'
    import { LIST_BLOG } from '../../../api/blog'

    const blogs = ref ([{total:0, items:[]}])
    const queryLoading = ref(false)
    const queryBlogs = async () => {
        try {
            const resp = await LIST_BLOG()
            blogs.value = resp
        }finally {
            queryLoading.value = false
        }

        

        // const resp = await LIST_BLOG()
        // console.log(resp)
    }

    // 页面渲染之前，需要把数据提前准备好
    onBeforeMount(async()=>{
        await queryBlogs()

    })

</script>
  
<style lang="css" scoped>


.table-op {
    height: 46px;
    display: flex;
    justify-content: space-between;
    align-items: center;
}

:deep(.arco-table .arco-spin) {
  height: unset;
}
</style>
  