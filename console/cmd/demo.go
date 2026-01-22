package cmd

import (
	"fmt"
	"time"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/spf13/cobra"
)

var demoCmd = &cobra.Command{
	Use:   "demo",
	Short: "命令行工具",
	Long:  `略`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(time.Now().UnixNano())
		fmt.Println([]byte(gconv.String(time.Now().UnixNano())))
		fmt.Println(time.Now().Format("15:04") >= "10:31")
	},
}
