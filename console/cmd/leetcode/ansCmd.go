package leetcode

import (
	"github.com/spf13/cobra"
)

var queNums int
var AnsCmd = &cobra.Command{
	Use:   "leetcode",
	Short: "leetcode",
	Long:  `略`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("demo called" + strconv.Itoa(queNums))
		switch queNums {
		case 1161:
			Q1161()

		}
	},
}

func init() {
	AnsCmd.PersistentFlags().IntVarP(&queNums, "nums", "n", 0, "leetcode nums")
}
