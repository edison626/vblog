import axios from "axios";
import { Message } from '@arco-design/web-vue'
//import { useRouter } from "vue-router";

var instance = axios.create({
    // 后端的URL地址， 沿用vite 的配置
    baseURL: '',
    //超时时间
    timeout: 5000,
    // 后端Gin 使用的Bind函数, 而非BindJson, 补充请求Data是哪个格式
    headers: {'Content-Type': 'application/json'}

});

// 通过响应拦截器统一处理异常
instance.interceptors.response.use(
    (resp) => {
        return resp
    },
    (error) => {
        //const router = useRouter()
        let msg = error.message

        //处理自定义异常
        if (error.response.data && error.response.data.message){
            // 通用逻辑处理
            msg = error.response.data.message

            //自定义业务逻辑处理
            switch (error.response.data.message){
                // token 过期,跳转到Login页面
                case 5001:
                    window.location.assign('/login')
                    break;
                default:
                    break;
            }
            // 是否要注入 Error，业务页面需要拿到异常
            // 只需要处理业务异常
            return Promise.reject(error.response.data)
        }

        // 直接吧异常信息显示出来
        Message.error(`${msg}`)


    }
)


export default instance 