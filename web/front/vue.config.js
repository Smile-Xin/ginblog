module.exports = {
  transpileDependencies: ['vuetify'],
  assetsDir: 'static',
  publicPath:"./",
  chainWebpack: config => {
    config.plugin('html').tap(args => {
      args[0].title = '欢迎来到GinBlog'
      return args
    })
  },
  productionSourceMap: false
}
