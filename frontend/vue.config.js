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
}