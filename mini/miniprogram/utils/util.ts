import { CallApiReq, HOST_NAME } from "../mdoel/model"
import { hideLoading, showLoading } from "./ui"

export const formatTime = (date: Date) => {
  const year = date.getFullYear()
  const month = date.getMonth() + 1
  const day = date.getDate()
  const hour = date.getHours()
  const minute = date.getMinutes()
  const second = date.getSeconds()

  return (
    [year, month, day].map(formatNumber).join('/') +
    ' ' +
    [hour, minute, second].map(formatNumber).join(':')
  )
}

const formatNumber = (n: number) => {
  const s = n.toString()
  return s[1] ? s : '0' + s
}



export function SendGet(url: string) {
  let req: CallApiReq = {
    url: url,
    method: "GET",
    message: "处理中"
  }
  return CallApi(req)
}

export function SendPost<T>(url: string, data: any) {
  let req: CallApiReq = {
    url: url,
    method: "POST",
    data: data,
    message: "处理中",
    showLoading: true
  }
  return CallApi<T>(req)
}

export function CallApi<T>(obj: CallApiReq) {
  return new Promise<T>(function (resolve, reject) {
    if (obj.showLoading) {
      showLoading(obj.message ? obj.message : '加载中...');
    }
    var data = {};
    if (obj.data) {
      data = obj.data;
    }
    var contentType = 'application/json';
    if (obj.contentType) {
      contentType = obj.contentType;
    }

    var method = 'GET';
    if (obj.method) {
      method = obj.method;
    }

    wx.request({
      url: HOST_NAME + obj.url,
      data: data,
      method: method,
      //添加请求头
      header: {
        'Content-Type': contentType,
        'token': wx.getStorageSync('token') //获取保存的token
      },
      //请求成功
      success: function (res) {
        console.log('===============================================================================================')
        console.log('==    接口地址：' + obj.url);
        console.log('==    接口参数：' + JSON.stringify(data));
        console.log('==    请求类型：' + method);
        console.log("==    接口状态：" + res.statusCode);
        console.log("==    接口数据：" + JSON.stringify(res.data));
        console.log('===============================================================================================')
        if (res.statusCode == 200) {
          let resData = res.data as T
          resolve(resData);
        } else if (res.statusCode == 401) {//授权失效
          reject("登录已过期");
        } else {
          //请求失败
          reject("请求失败：" + res.statusCode)
        }
      },
      fail: function (err) {
        //服务器连接异常
        console.log('===============================================================================================')
        console.log('==    接口地址：' + url)
        console.log('==    接口参数：' + JSON.stringify(data))
        console.log('==    请求类型：' + method)
        console.log("==    服务器连接异常")
        console.log('===============================================================================================')
        reject("服务器连接异常，请检查网络再试");
      },
      complete: function () {
        hideLoading();
      }
    })
  });
}

export function GenImageUrl(imageNo: string) {
  return `${HOST_NAME}/assets/${imageNo}?token=${wx.getStorageSync('token')}`
}

export function IsBlank(st: string) {
  if (st == undefined || st == null || st == '' || st.length <= 0) {
    return true
  }
}
