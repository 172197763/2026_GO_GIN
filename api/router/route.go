package router

import (
	"gin_test/api/app/controllers"
	"gin_test/api/database"
	"gin_test/api/router/middleware"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	defaultRoute(r)                                       //默认路由、不会触发后续的注册的全局中间件
	r.Use(middleware.Log())                               //全局中间件
	r.GET("/index", middleware.Test(), controllers.Index) //全局路由+中间件 会触发log中间件
	userRoute(r)                                          //用户路由 会触发log中间件

	r.LoadHTMLGlob(filepath.Join(database.GetRootDir(), "templates/*"))
	r.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(404, "404.html", gin.H{
			"title": "你迷路啦",
			"path":  ctx.Request.URL.Path,
		})
	})
}
