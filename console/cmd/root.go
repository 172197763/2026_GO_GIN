package cmd

import (
	"gin_test/console/cmd/leetcode"

	"github.com/spf13/cobra"
)

// 创建根命令
var RootCmd = &cobra.Command{
	Use:   "cli",
	Short: "命令行工具",
	Long:  `略`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	RootCmd.AddCommand(demoCmd)
	RootCmd.AddCommand(timeCmd)
	RootCmd.AddCommand(leetcode.AnsCmd)
}
