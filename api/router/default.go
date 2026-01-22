package router

import (
	"gin_test/api/app/controllers/upload"

	"github.com/gin-gonic/gin"
)

func defaultRoute(r *gin.Engine) {
	r.POST("/uploadFile", (&upload.Upload{}).UploadFile)
}
