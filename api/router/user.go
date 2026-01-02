package router

import (
	"gin_test/api/app/controllers/user"

	"github.com/gin-gonic/gin"
)

func userRoute(r *gin.Engine) {
	//定义路由组
	def := r.Group("/user")
	//路由关联方法
	def.GET("/name", user.NewUser().MyName)
	def.GET("/sex", user.NewUser().MySex)
	def.GET("/updatebatch", user.NewUser().UpdateBatch)
	def.POST("/reg", user.NewUser().Reg)
	def.GET("/sendphonemsg", user.NewUser().SendPhoneMsg)
	def.GET("/test", user.NewUser().Test)
}
