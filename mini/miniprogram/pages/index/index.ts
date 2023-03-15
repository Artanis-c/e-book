// index.ts

import { ApiResponse, BookCategory, BookListMode, BookListReq, PageModel } from "../../mdoel/model"
import { GenImageUrl, SendGet, SendPost } from "../../utils/util"

// 获取应用实例
const app = getApp<IAppOption>();
var BookList: Array<BookListMode> = [];
var BookCategoryList:Array<BookCategory>=[];
Page({
  data: {
    bookList: BookList,
    pageIndex: 1,
    pageSize: 10,
    total: 0,
    barCode: '',
    bookName: '',
    categoryList:BookCategoryList
  },
  // 事件处理函数
  bindViewTap() {
    wx.navigateTo({
      url: '../logs/logs',
    })
  },
  onLoad() {
    wx.setNavigationBarTitle({
      title: '我的书架'
    })
    this.lunchData();
    this.queryCategory();
  },
  onShow() {
    this.getTabBar().setData({ selected: 0 })
  },
  onAdd() {
    wx.navigateTo({
      url: '/pages/book/book'
    })
  },
  lunchData() {
    let req: BookListReq = {
      pageIndex: this.data.pageIndex,
      pageSize: this.data.pageSize,
      barCode: this.data.barCode,
      bookName: this.data.bookName
    }
    let callBack = SendPost<ApiResponse<PageModel<BookListMode>>>('/api/book/list', req)
    callBack.then(res => {
      let resData = res.data.records
      let oldData = this.data.bookList
      resData.forEach(x => {
       x.image= GenImageUrl(x.image)
        oldData.push(x)
      })
      this.setData({ bookList: oldData, pageIndex: res.data.pageIndex, pageSize: res.data.pageSize, total: res.data.total })

    })
  },
  queryCategory(){
    let callBack = SendPost<ApiResponse<Array<BookCategory>>>("/api/category/list", {})
    callBack.then(res => {
      this.setData({ categoryList: res.data })
    })
  }
})
