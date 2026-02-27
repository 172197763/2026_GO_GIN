package cmd

import (
	"gin_test/api/common"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

// interface相关知识
var syncCondCmd = &cobra.Command{
	Use:   "syncCond",
	Short: "sync.Cond简单示例",
	Long:  `略`,
	Run: func(cmd *cobra.Command, args []string) {
		syncCondHandle()
	},
}
var condFlat = false

func syncCondHandle() {
	common.PrintT("sync.Cond演示start")
	var cond = sync.NewCond(&sync.Mutex{})
	go condRead("消费者1", cond)
	go condRead("消费者2", cond)
	go condRead("消费者3", cond)
	time.Sleep(time.Second)
	condWrite(cond)
	time.Sleep(time.Second)
	common.PrintT("sync.Cond演示end")
}

// 消费者
func condRead(name string, cond *sync.Cond) {
	cond.L.Lock()
	defer cond.L.Unlock()
	for !condFlat {
		//调用 Wait 会自动释放锁 c.L，并挂起调用者所在的 goroutine
		//如果其他协程调用了Signal或Broadcast唤醒了该协程，那么Wait方法在结束阻塞时，会重新给c.L加锁，并且继续执行 Wait 后面的代码。
		cond.Wait()
	}
	common.PrintT("%s消费完成", name)
}

// 生产者
func condWrite(cond *sync.Cond) {
	cond.L.Lock()
	defer cond.L.Unlock()
	condFlat = true
	common.PrintT("生产完成")
	//唤醒所有等待的goroutine
	cond.Broadcast()
	//唤醒单个等待的goroutine
	// cond.Signal()
}
