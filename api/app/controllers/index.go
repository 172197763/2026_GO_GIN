package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	fmt.Println("这里是首页index")
	ctx.JSON(200, gin.H{
		"path": "index",
	})
}
