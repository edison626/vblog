<template>
    <div>博客编辑</div>
    <div>{{ router.currentRoute.value }}</div>
</template>
  
<script setup>
import { useRouter } from 'vue-router'

const router = useRouter()

router.currentRoute.value

</script>
  
  <style lang="css" scoped></style>