// 实例化axios 封装相关功能
import axios from 'axios'
// 替换为 Element Plus 的 ElMessage
// import { ElMessage } from 'element-plus'
// import store from '@/store'
import router from '@/router'
// import { getTimeStamp } from '@/utils/auth'

const TimeOut = 3600 // 定义token有限期 1小时

// 创建axios实例
const service = axios.create({
  baseURL: '', // 设置axios请求的基础的基础地址
  // baseURL: process.env.VUE_APP_BASE_API || '127.0.0.1:8888', // 设置axios请求的基础的基础地址
  timeout: 5000
})
// 请求拦截器
service.interceptors.request.use(
  // config 是请求的配置信息
  config => {
    // 需要统一的去注入token
    // if (store.getters.token) {
    //   if (IsCheckTimeOut()) {
    //     // 如果它为true表示 过期了
    //     // token没用了 因为超时了
    //     store.dispatch('user/logout') // 登出操作
    //     // 跳转到登录页面
    //     router.push('/login')
    //     // 直接抛出错误，跳出逻辑
    //     return Promise.reject(new Error('The token timeout'))
    //   }
    //   // 如果token存在 注入token
    //   config.headers['Authorization'] = `Bearer ${store.getters.token}`
    // }
    return config // 必须返回配置
  },
  error => {
    // 如果报错，原样返回错误
    return Promise.reject(error)
  }
)
// 响应拦截器
service.interceptors.response.use(
  response => {
    // 响应拦截器的第一个参数，是一个函数，是响应成功的回调
    // const { data } = response.data
    // if (success) {
    //   // return data
    //   return response
    // } else {
    //   // 使用 ElMessage 提示错误消息
    //   // ElMessage.error(message)
    //   return Promise.reject(new Error(message))
    // }
    return response.data
  },
  error => {
    // // 响应拦截器的第二个参数，是一个函数，是响应失败的回调
    // if (
    //   error.response &&
    //   error.response.data &&
    //   error.response.data.code === 10002
    // ) {
    //   // 当等于10002的时候 表示 后端告诉我token超时了
    //   store.dispatch('user/logout') // 登出action 删除token
    //   router.push('/login')
    // } else {
    //   // 使用 ElMessage 提示错误消息
    //   ElMessage.error(error.message)
    // }
    return Promise.reject(new Error(error.message)) // 返回执行错误 让当前的执行链跳出成功 直接进入 catch
  }
)

// 是否超时
// // 超时逻辑 (当前时间  - 缓存中的时间) 是否大于 时间差
// function IsCheckTimeOut (params) {
//   var currentTime = Date.now() // 当前时间戳
//   var timeStamp = getTimeStamp() // 缓存时间戳
//   return (currentTime - timeStamp) / 1000 > TimeOut
// }

export default service // 导出axios 实例