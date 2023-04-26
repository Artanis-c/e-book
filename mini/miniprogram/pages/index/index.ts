// index.ts

import { ApiResponse, BookCategory, BookListMode, BookListReq, PageModel } from "../../mdoel/model"
import { GenImageUrl, SendGet, SendPost } from "../../utils/util"

// 获取应用实例
const app = getApp<IAppOption>();
var BookList: Array<BookListMode> = [];
var BookCategoryList: Array<BookCategory> = [];
Page({
  data: {
    bookList: BookList,
    pageIndex: 1,
    pageSize: 10,
    total: 0,
    barCode: '',
    bookName: '',
    categoryNo: '',
    categoryList: BookCategoryList
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

  },
  onShow() {
    this.lunchData();
    this.queryCategory();
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
      bookName: this.data.bookName,
      categoryNo: this.data.categoryNo
    }
    let callBack = SendPost<ApiResponse<PageModel<BookListMode>>>('/api/book/list', req)
    callBack.then(res => {
      let resData = res.data.records
      let oldData = this.data.bookList
      resData.forEach(x => {
        x.image = GenImageUrl(x.image)
        oldData.push(x)
      })
      this.setData({ bookList: oldData, pageIndex: res.data.pageIndex, pageSize: res.data.pageSize, total: res.data.total })
    })
  },
  queryCategory() {
    let callBack = SendPost<ApiResponse<Array<BookCategory>>>("/api/category/list", {})
    callBack.then(res => {
      this.setData({ categoryList: res.data })
    })
    console.log(this.data.categoryList)
  },
  onLower(e: any) {
    let oldPageIndex = this.data.pageIndex
    let newPageIndex = oldPageIndex + 1;
    this.setData({ pageIndex: newPageIndex })
    this.lunchData()
  }, onCategoryChange(event: any) {
    let category = this.data.categoryList[event.detail]
    this.setData({ pageIndex: 1, categoryNo: category.categoryNo })
    this.queryData(category.categoryNo,'','')
  }, onSearch(event:any){
    this.setData({bookName:event.detail})
    this.queryData('','',event.detail)
  },
  onScan() {
    var that = this
    wx.scanCode({
      onlyFromCamera: true,
      success(res) {
        that.setData({ bookName: res.result })
      }
    })
  }
  ,queryData(categoryNo: string, barCode: string, bookName: string) {
    let req: BookListReq = {
      pageIndex: 1,
      pageSize: this.data.pageSize,
    }
    if(barCode!=''){
      req.barCode=barCode
    }
    if(bookName!=''){
      req.bookName=bookName
    }
    if(categoryNo!=''){
      req.categoryNo=categoryNo
    }
    let callBack = SendPost<ApiResponse<PageModel<BookListMode>>>('/api/book/list', req)
    callBack.then(res => {
      let resData = res.data.records
      let newData : Array<BookListMode>=[]
      resData.forEach(x => {
        x.image = GenImageUrl(x.image)
        newData.push(x)
      })
      this.setData({ bookList: newData, pageIndex: res.data.pageIndex, pageSize: res.data.pageSize, total: res.data.total })
    })
  }
})
