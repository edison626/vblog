import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    // 登陆页
    {
      path: '/login',
      name: 'Login',

      component: () => import('../views/login/LoginView.vue')
    },
    // 前台
    {
      path: '/frontend',
      name: 'FrontendLayout',
      component: () => import('../views/frontend/LayoutView.vue'),
      children: [
        {
          // blogs --> /frontend/blogs
          path: 'blogs',
          name: 'FrontendBlogs',
          component: () => import('../views/frontend/blog/ListView.vue')
        }
      ]
    },
    //后台
    {
      path: '/backend',
      name: 'BackendLayout',
      component: () => import('../views/backend/LayoutView.vue'),
      children: [
        {
          path: 'blogs',
          name: 'BackendBlogs',
          component: () => import('../views/backend/blog/ListView.vue')
        },
        {
          path: 'blogs_edit',
          name: 'BackendEdit',
          component: () => import('../views/backend/blog/EditView.vue')
        },
        {
          path: 'comments',
          name: 'CommentList',
          component: () => import('../views/backend/comment/ListPage.vue')
        }
      ]
    }
  ]
})

export default router
