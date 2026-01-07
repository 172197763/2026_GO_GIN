package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var demoCmd = &cobra.Command{
	Use:   "demo",
	Short: "命令行工具",
	Long:  `略`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("demo called")
	},
}
