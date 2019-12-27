## 微微博（因为功能太少所以是比微博还微小的博客）

### 小组名单
```
17343092
17343093
17343098
17343099
173430101
173430102
```

---

### 运行方式
```bash
# 使用代理下载go mod 依赖，bash下运行
GOPROXY="https://goproxy.io" go mod download
# 构建前端
cd frontend
npm install --regstry=https://registry.npm.taobao.org
npm run build
# 运行服务端程序
cd ..
go run main.go
```

### 开发指北
* 和数据库有关的函数放在`micro-microblog/database`包里面
* 路由写在`micro-microblog/router`包里面
* 最好写一些单元测试，在前端还没有对接之前只有单元测试能让人放心
* 因为是 go module，所以依赖的包例如[gin](https://www.github.com/go-gonic/gin)、[bolt](https://www.github.com/boltdb/bolt)、[uuid](https://github.com/satori/go.uuid) 会在构建运行之前自动安装，中间可能会遇到天朝特色的安装失败，可以百度一下`go-proxy`关键字，配置一下包的代理
* 已经安装的包和 Go 的版本在`./go.mod`文件里，可以自行查询并安装别的包（如有需要）
* clone 下来后，建议先建一个个人分支，名字可以为`dev/你的用户名`，然后每次开发前，都 pull 一下 master 分支。开发完成一个阶段（比如说写完一个接口）后，就 merge 到 master 分支里
* 写接口的同学将所需的跟数据库有关的函数的签名放进`./database/database.go`目录下（记得在函数签名前加上注释说明该函数的功能），负责和数据库对接的同学再进行实现
* 还有其余的备忘内容也写在这里