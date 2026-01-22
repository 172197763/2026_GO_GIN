package user

import (
	"fmt"
	"gin_test/api/app/services/dict"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

func (u *User) Index(ctx *gin.Context) {
	str := "i am user[index]"
	fmt.Println(str)
	ctx.String(http.StatusOK, str)
}
func (u *User) Test(ctx *gin.Context) {
	content := fmt.Sprintf("有新的门店审核待处理。门店名称:%v,门店ID:%v,用户ID:%v,时间:%v", "测试", "测试1", "测试2", time.Now().Format("2006-01-02 15:04:05"))
	mapSafe("key1", "value1")
	ctx.JSON(http.StatusOK, gin.H{"data2": 12, "msg": content})
}
func mapSafe(k string, v string) {
	m := dict.NewShopTypeMap()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			m.Set(fmt.Sprintf("%d", n), k+v)
		}(i)
	}
	wg.Wait()
	for key, value := range *m.GetData() {
		fmt.Println("mapdata:key:", key, "value:", value)
	}
	fmt.Println(m.Len())
}

func mapUnsafe() {
	m := make(map[int]int)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			m[n] = n // 并发写 -> 不安全
		}(i)
	}
	wg.Wait()
	fmt.Println(len(m))
}
