import { HOST_NAME } from "./mdoel/model"

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
        wx.request({
          url: `${HOST_NAME}/api/user/login`,
          data: data,
          method:'POST',
          success:res=>{
            console.log(res)
          }
        })
      },
    })
  },
})