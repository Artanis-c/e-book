Component({
  data: {
    selected: 0,
    color: "#7A7E83",
    selectedColor: "#3cc51f",
    list: [{
      pagePath: "/pages/index/index",
      iconName: "orders-o",
      text: "菜单"
    }, {
      pagePath: "/pages/my/my",
      iconName: "user-o",
      text: "我的"
    }]
  },
  lifetimes: {},
  pageLifetimes: {},
  attached() {},
  methods: {
    switchTab(e) {
      const data = e.currentTarget.dataset
      const url = data.path
      wx.switchTab({
        url: url,
      })
    }
  }
})