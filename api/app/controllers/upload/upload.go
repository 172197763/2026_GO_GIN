package upload

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Upload struct {
}

// 上传单文件
func (u *Upload) Index(ctx *gin.Context) {
	str := "i am upload[index]"
	fmt.Println(str)
	ctx.String(http.StatusOK, str)
}

// 上传单文件
func (u *Upload) UploadFile(ctx *gin.Context) {
	// 单文件
	file, _ := ctx.FormFile("file")
	log.Println(file.Filename)

	dst := "./upload/file" + file.Filename
	// 上传文件至指定的完整文件路径
	ctx.SaveUploadedFile(file, dst)

	ctx.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
