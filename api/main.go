// main.go
package main

import (
	"gin_test/api/app/services/redis_tool"
	"gin_test/api/database"
	"gin_test/api/router"
	_ "net/http/pprof"

	"github.com/gin-gonic/gin"
)

var redis *redis_tool.RedisStreamListener

func init() {
	// log.Println(http.ListenAndServe("localhost:6060", nil))
	err := database.InitDatabase()
	if err != nil {
		panic(err) // 或者使用 log.Fatal(err)
	}
	redis = redis_tool.NewRedisStreamListener()

	// 启动监听
	redis.StartListening("mystream", "mygroup", "consumer1")
}
func main() {
	defer redis.Stop()
	// 创建默认的 Gin 路由引擎
	r := gin.Default()
	//初始化路由数据
	router.Init(r)
	// 启动 HTTP 服务器，默认监听 :8080
	r.Run(":8080")
}
