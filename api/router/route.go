package router

import (
	"gin_test/api/database"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	//调用路由
	userRoute(r)
	defaultRoute(r)
	r.LoadHTMLGlob(filepath.Join(database.GetRootDir(), "templates/*"))
	r.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(404, "404.html", gin.H{
			"title": "你迷路啦",
			"path":  ctx.Request.URL.Path,
		})
	})
}
