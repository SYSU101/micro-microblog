package router

import (
	"io"
	"log"
	"micro-microblog/middlewares"
	"os"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine = nil

// Route 初始化路由并监听 address 指定的地址
func Route(address string) {
	// 初始化日志文件
	gin.DisableConsoleColor()
	logFile, err := os.Create("./logs/gin.log")
	if err != nil {
		log.Fatal(err)
		return
	}
	gin.DefaultWriter = io.MultiWriter(logFile)
	router = gin.Default()

	// 各种全局的中间件写在这里
	router.Use(middlewares.GetSessionID())

	// 各种路由写在这里，要确保 routerStatic 是最后执行的
	routeStatic()

	// 监听
	router.Run(address)
}
