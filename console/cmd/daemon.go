package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

// 守护进程模式启动进程
var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "守护进程",
	Long:  `守护进程`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("支付成功unipush推送--start")
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// 监听系统信号
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

		go func() {
			//业务逻辑
		}()
		for {
			select {
			case sig := <-sigChan:
				fmt.Printf("\n收到信号 %v，正在关闭服务...\n", sig)
				return
			case <-ctx.Done():
				fmt.Println("支付成功unipush推送--end")
				return
			}
		}
	},
}
