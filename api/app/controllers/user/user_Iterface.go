package user

import "github.com/gin-gonic/gin"

type IUser interface {
	MyName(ctx *gin.Context)
	MySex(ctx *gin.Context)
	UpdateBatch(ctx *gin.Context)
	SendPhoneMsg(ctx *gin.Context)
	Reg(ctx *gin.Context)
	Test(ctx *gin.Context)
	Index(ctx *gin.Context)
}
