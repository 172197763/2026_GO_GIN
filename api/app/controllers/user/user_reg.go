package user

import (
	"context"
	"log"
	"net/http"
	"time"

	userService "gin_test/api/app/services/user"
	"gin_test/api/common"
	"gin_test/proto/calc"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type RegRequest struct {
	Phone string `json:"phone" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Age   int    `json:"age"`
	Sex   *int   `json:"sex"`
}

// 注册&调用rpc
func (u *User) Reg(ctx *gin.Context) {
	req_data := &RegRequest{}
	err := common.GetParams(ctx, &req_data)
	if err != nil {
		return
	}
	rpcRes, err := u.rpcFun(1, 2)
	if err != nil {
		rpcRes = -1
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  req_data,
		"data": rpcRes,
	})
}
func (u *User) rpcFun(num1 int32, num2 int32) (int32, error) {
	// ✅ 使用 NewClient 替代 Dial
	client, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// grpc.WithDefaultCallOptions(grpc.WaitForReady(true)), // 可选：等待服务就绪
	)
	if err != nil {
		log.Printf("Failed to create client: %v", err)
		return 0, err
	}
	// 读取项目下 ./file/* 的文件（需要在文件顶部导入 "os" 和 "path/filepath"）

	calculator := calc.NewCalculatorClient(client)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	resp, err := calculator.Add(ctx, &calc.AddRequest{A: num1, B: num2})
	if err != nil {
		log.Printf("Add failed: %v", err)
	}

	log.Printf("Result: %d", resp.GetResult())
	return resp.GetResult(), nil
}

// 协程体验
func (u *User) SendPhoneMsg(ctx *gin.Context) {
	phone := "13000000001"
	resChan := make(chan string, 2)
	// 确保 channel 最终会被关闭
	defer close(resChan)
	go func(phone string) {
		resChan <- userService.SendPhoneMsg("13000000001")
	}(phone)
	go func() {
		resChan <- userService.Reg()
	}()
	//不确保返回结果顺序
	res1 := <-resChan
	res2 := <-resChan
	close(resChan)

	ctx.JSON(http.StatusOK, gin.H{"data2": res1 + res2})
}
