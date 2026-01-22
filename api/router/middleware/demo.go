package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Test() gin.HandlerFunc {
	return func(c *gin.Context) {
		str := "[全局中间件Test]"
		// 🟢 阶段1：Handler 执行前（前置逻辑）
		fmt.Println(str + "执行Before:")

		// ⚡ 关键：调用 c.Next() → 执行下一个中间件或最终 Handler
		c.Next()
		// 不继续执行后续步骤
		// c.Abort()

		// 🔴 阶段2：Handler 执行后（后置逻辑）
		fmt.Println(str + "执行After:")
	}
}
func Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		str := "[日志中间件Log]"
		// 🟢 阶段1：Handler 执行前（前置逻辑）
		fmt.Println(str + "执行Before:" + c.FullPath())

		// ⚡ 关键：调用 c.Next() → 执行下一个中间件或最终 Handler
		c.Next()
		// 不继续执行后续步骤
		// c.Abort()

		// 🔴 阶段2：Handler 执行后（后置逻辑）
		fmt.Println(str + "执行After:" + c.FullPath())
	}
}
