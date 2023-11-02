// 保持我们当前的应用程序的状态
import { useStorage } from '@vueuse/core'

export const state = useStorage(
    'vblog',
    { 
        is_login: false,
        token: {},
    },
    localStorage,
    { mergeDefaults: true },
)

//
// console.log(state.value.hello) // 'nihao', from storage
// console.log(state.value.greeting) // 'hello', from merged default value