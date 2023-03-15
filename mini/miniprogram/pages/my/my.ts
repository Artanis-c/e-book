// pages/my/my.ts

import { GridItemModel } from "miniprogram/mdoel/model"
var gridItem:Array<GridItemModel> =[
  {
    name:'图书分类',
    icon:'apps-o',
    page:'/pages/type/type'
  },
  {
    name:'图书管理',
    icon:'newspaper-o',
    page:'/pages/book/book'
  }
]

Page({


  /**
   * 页面的初始数据
   */
  data: {
    gridItem:gridItem
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad() {
 
  },

  /**
   * 生命周期函数--监听页面初次渲染完成
   */
  onReady() {

  },

  /**
   * 生命周期函数--监听页面显示
   */
  onShow() {
    let tab= this.getTabBar()
    tab.setData({ selected: 1})
  },

  /**
   * 生命周期函数--监听页面隐藏
   */
  onHide() {

  },

  /**
   * 生命周期函数--监听页面卸载
   */
  onUnload() {

  },

  /**
   * 页面相关事件处理函数--监听用户下拉动作
   */
  onPullDownRefresh() {

  },

  /**
   * 页面上拉触底事件的处理函数
   */
  onReachBottom() {

  },

  /**
   * 用户点击右上角分享
   */
  onShareAppMessage() {

  }
})