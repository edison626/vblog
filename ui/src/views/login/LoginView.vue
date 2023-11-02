<template>
    <div class="login-page">
        <div class="login-form">
            <a-form-item hide-label >
                <span class="login-title"> 欢迎登陆博客系统 </span>
            </a-form-item>
            <a-form :model="loginForm"  @submit="handleSubmit">
                <!-- 定义里面哪个字段 -->
                <a-form-item hide-label field="username" :rules="[{required:true,message:'请输入用户名'}]">
                    <a-input v-model="loginForm.username" placeholder="请输入用户名" allow-clear>
                        <template #prefix>
                            <icon-user />
                        </template>
                    </a-input> 
                </a-form-item>
                <a-form-item hide-label field="password" :rules="[
                    {required:true,message:'请输入密码'},
                    {minLength:5,message:'密码至少6位数'}]">
                    <a-input-password v-model="loginForm.password" placeholder="请输入密码">
                        <template #prefix>
                            <icon-lock />
                        </template>
                    </a-input-password> 
                </a-form-item>
                <a-form-item hide-label>
                <a-button html-type="submit" style="width: 100%;" type="primary" >登陆</a-button>
                </a-form-item>
            </a-form>
        </div>
    </div>
</template>

<script setup>
import {ref} from 'vue'
import { LOGIN } from '../../api/token'
import { useRouter } from 'vue-router';
import { state } from '../../stores/app';

const router = useRouter()

//定义表单数据提交
// 表单对应的响应数据
// 按照Vblog login API 来设计
const loginForm = ref ({
    username: '',
    password: ''
})

//表单数据提交函数
const handleSubmit = async (data) => {
    
    // 表单校验成功，才与后端交互
    if(data.errors !== undefined ) {
        return
    }

    try {
        const resp = await LOGIN(loginForm.value)
        //console.log(resp)

        //保留登陆状态
        state.value.is_login = true;
        state.value.token = resp.data

        // 需要进行跳转， vue router 的 Router 对象
        // 通过vue router 库来获取一个router对象
        router.push({name: 'BackendBlogs'})
        
    } catch (error) {
        // 这就是promise 的Reject error
        console.log(error);
    }
}
</script>

<style lang="css" scoped>

.login-page {
    height: 100vh;
    width: 100vw;
    display: flex;
    align-items: center;
    justify-content: center;
}

.login-form {
    
    display: flex;
    align-items: center;
    flex-direction: column;
    height: 400px;
    width: 400px;

}
.login-title {
    display: flex;
    width: 100%;
    justify-content: center;
}
</style>