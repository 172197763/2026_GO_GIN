// main.go
package main

import (
	"fmt"
	apiConfig "gin_test/api/config"
	"gin_test/api/router"
	_ "net/http/pprof"

	"github.com/gin-gonic/gin"
)

// var redis *redis_tool.RedisStreamListener

func init() {
}
func main() {
	// defer redis.Stop()
	// 创建默认的 Gin 路由引擎
	r := gin.Default()
	//初始化路由数据
	router.Init(r)
	//读取配置端口号
	port, err := apiConfig.GetConfigString("APP.PORT")
	if err != nil {
		panic(fmt.Errorf("%w:APP.PORT", err))
	}
	env, err := apiConfig.GetConfigString("APP.ENV")
	if err != nil {
		panic(fmt.Errorf("%w:APP.ENV", err))
	}
	fmt.Printf("服务启动，监听端口为[%s]，项目环境为[%s]\n", port, env)
	// 启动 HTTP 服务器
	r.Run(":" + port)
}
