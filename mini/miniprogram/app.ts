import { ApiResponse, CallApiReq, HOST_NAME } from "./mdoel/model"
import { CallApi, SendGet, SendPost } from "./utils/util"

// app.ts
App<IAppOption>({
  globalData: {},
  onLaunch() {
    // 展示本地存储能力
    const logs = wx.getStorageSync('logs') || []
    logs.unshift(Date.now())
    wx.setStorageSync('logs', logs)
    // 登录
    wx.login({
      success: data => {
        let res = SendPost<ApiResponse<string>>("/api/user/login", data)
        res.then(res => {
          wx.setStorageSync('token', res.data)
        })
      },
    })
  },
})