import client from './client'
// 区分API请求, 可以全大写
// config?: AxiosRequestConfig<D>
// Get 请求, params?   url?a=1&b2
// 区分api 请求 ，可以全部大写
export var LIST_BLOG = (params) => client.get('/api/vblog/v1/blogs/', {params})

