package router

import (
	"gin_test/api/app/controllers/upload"

	"github.com/gin-gonic/gin"
)

func defaultRoute(r *gin.Engine) {
	r.POST("/uploadFile", (&upload.Upload{}).UploadFile)
	// r.GET("/index", (&upload.Upload{}).Index)//会报错，因为全局已注册该路由
	r.GET("/upload/index", (&upload.Upload{}).Index)
}
