const path = require('path');
const apiMocker = require('mocker-api');

module.exports = {
  assetsDir: './assets',
  configureWebpack: {
    devServer: {
      before(app){
        apiMocker(app, path.resolve('./mock/index.js'))
      }
    },
  },
  css: {
    loaderOptions: {
      less: {
        modifyVars: {
          'primary-color': '#060810',
        },
        javascriptEnabled: true,
      },
    },
  },
}