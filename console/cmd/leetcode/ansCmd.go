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
		case 1339:
			Q1339()
		case 865:
			Q865()
		case 1126:
			Q1126()
		case 3637:
			Q3637()
		case 3640:
			Q3640()
		case 3379:
			Q3379()

		}
	},
}

func init() {
	AnsCmd.PersistentFlags().IntVarP(&queNums, "nums", "n", 0, "leetcode nums")
}
