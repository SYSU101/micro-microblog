package router

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine = nil

func Route() {
	// 执行一些初始化
	gin.DisableConsoleColor()
	logFile, err := os.Create("./logs/gin.log")
	if err != nil {
		log.Fatal(err)
		return
	}
	gin.DefaultWriter = io.MultiWriter(logFile)
	router = gin.Default()

	// 各种路由写在这里，要确保 routerStatic 是最后执行的
	routeStatic()

	router.Run(":8080")
}
