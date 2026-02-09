package cmd

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/mattn/go-runewidth"
	"github.com/spf13/cobra"
)

type data1 struct {
	Name string
	Sn   string
	Type int
}

var demoCmd = &cobra.Command{
	Use:   "demo",
	Short: "命令行工具",
	Long:  `略`,
	Run: func(cmd *cobra.Command, args []string) {
		for i := 390000; i < 410000; i++ {
			tmp := i
			cj(tmp)
		}
	},
}

// cj 根据上证指数6位数字作为seed，生成随机数，并将nameArr中的每个元素对应的随机数打印出来
func cj(n int) {
	seed := int64(n) //取周一3点收市时，上证指数6位数字作为seed
	nameArr := []string{"jvav攻城狮练习两年半", "摸鱼王小二", "臭咸鱼20299", "小鳄鱼crocodile", "小熊猫幺儿"}
	randSource := rand.New(rand.NewSource(seed/int64(len(nameArr)) + seed%int64(len(nameArr))))
	fmt.Printf("----------%d------------\n", n)
	for _, v := range nameArr {
		fmt.Printf("%s : %d\n", padToWidth(v, 20), randSource.Intn(100))
	}
}

func padToWidth(s string, targetWidth int) string {
	actualWidth := runewidth.StringWidth(s) // 获取真实显示宽度
	if actualWidth >= targetWidth {
		// 安全截断（按显示宽度）
		return runewidth.Truncate(s, targetWidth, "...")
	}
	return s + strings.Repeat(" ", targetWidth-actualWidth)
}
