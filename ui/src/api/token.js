import client from './client'

// 颁发token 
export var LOGIN = (data) => {
    return client ({
        url: '/api/vblog/v1/tokens/',
        method: 'post',
        data: data
    })
}