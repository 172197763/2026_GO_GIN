package cmd

import (
	"context"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// time定时器演示
var timeCmd = &cobra.Command{
	Use:   "timeCmd",
	Short: "time定时器演示",
	Long:  `time定时器演示`,
	Run: func(cmd *cobra.Command, args []string) {
		funTimer()
		funTimerStop()
		funTicker()
	},
}

// Timer&reset演示
func funTimer() {
	println(time.Now().Format("2006-01-02 15:04:05") + "|funTimer--start")
	//指定时间后同个channel发送数据
	timer := time.NewTimer(time.Second * 2)
	//通过reset重置定时器时间
	timer.Reset(time.Second)
	//启动一个goroutine等待定时器触发
	go func() {
		select {
		case <-timer.C:
			println(time.Now().Format("2006-01-02 15:04:05") + "|触发NewTimer")
		case <-time.After(time.Second * 2):
			println(time.Now().Format("2006-01-02 15:04:05") + "|NewTimer超时")
		}
	}()
	//指定时间后执行函数
	time.AfterFunc(2*time.Second, func() {
		println(time.Now().Format("2006-01-02 15:04:05") + "|触发AfterFunc")
	})
	// 主goroutine等待3秒，确保看到定时器触发的输出
	time.Sleep(time.Second * 3)
	println(time.Now().Format("2006-01-02 15:04:05") + "|funTimer--end")
}

// Timer Stop演示
func funTimerStop() {
	println(time.Now().Format("2006-01-02 15:04:05") + "|funTimerStop--start")
	//指定时间后同个channel发送数据
	timer := time.NewTimer(time.Second * 2)
	ctx, cancel := context.WithCancel(context.Background())
	//启动一个goroutine等待定时器触发
	go func() {
		select {
		case <-timer.C:
			println(time.Now().Format("2006-01-02 15:04:05") + "|触发NewTimer")
		case <-ctx.Done():
			timer.Stop()
			println(time.Now().Format("2006-01-02 15:04:05") + "|NewTimer被停止")
		}
	}()
	cancel()
	time.Sleep(time.Second * 2)
	println(time.Now().Format("2006-01-02 15:04:05") + "|funTimerStop--end")
}

// Ticker演示
// @desc:Reset&Stop方法也适用于Ticker，Reset将会重置周期间隔
func funTicker() {
	logTitle := "[funTicker]"
	println(time.Now().Format("2006-01-02 15:04:05") + "|" + logTitle + "--start")
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*8)
	defer cancelFunc()
	go func() {
		var i int
		for {
			select {
			case <-ticker.C:
				println(time.Now().Format("2006-01-02 15:04:05") + "|" + logTitle + "触发Ticker" + strconv.Itoa(i))
				i++
			case <-timeout.Done():
				println(time.Now().Format("2006-01-02 15:04:05") + "|" + logTitle + "Ticker被终止")
				return
			}
		}
	}()
	time.Sleep(time.Second * 10)
	println(time.Now().Format("2006-01-02 15:04:05") + "|" + logTitle + "--end")
}
