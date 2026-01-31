package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

// 守护进程模式启动进程
var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "守护进程",
	Long:  `守护进程`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("daemon服务--start")
		ctx, cancel := context.WithCancel(context.Background())
		defer func() {
			fmt.Println("触发defer")
			cancel()
		}()

		// 监听系统信号
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

		go func() {
			doSomething(ctx)
		}()
		for {
			select {
			case sig := <-sigChan:
				fmt.Printf("\n收到信号 %v，正在关闭服务...\n", sig)
				return
			case <-ctx.Done():
				fmt.Println("daemon服务--end")
				return
			}
		}
	},
}

func doSomething(ctx context.Context) {
	ticker := time.NewTicker(time.Second * 3)
	for {
		select {
		case <-ticker.C:
			fmt.Println("start to do something")
			time.Sleep(3 * time.Second)
			fmt.Println("end to do something")
		case <-ctx.Done():
			fmt.Println("中断 to do something")
			return
		}
	}
}
