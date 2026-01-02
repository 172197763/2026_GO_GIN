package common

import (
	"github.com/gin-gonic/gin"
)

// 获取&检测参数
func GetParams(ctx *gin.Context, data interface{}) error {
	// 自动识别 Content-Type 并绑定
	if err := ctx.ShouldBind(data); err != nil {
		// panic("get params error:" + err.Error() + ctx.ClientIP())
		// ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		// 	"code": 400,
		// 	"msg":  err.Error(),
		// })
		ctx.JSON(400, gin.H{"error": err.Error()})
		ctx.Abort() // 中断请求处理
		return err
		// ctx.Next()
	}
	return nil
}
