package cmd

import (
	"fmt"
	"gin_test/api/common"
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
		// fmt.Println("2026-02-09" < "2026-02-10")
		// cj(412309)
		// fmt.Println(config.Get("mysql.host", "127.0.0.1"))
		// fmt.Println(config.Get("mysql.port", 3306))
		// mq, err := gmqtt.GetMQTTClient()
		// if err != nil {
		// 	fmt.Println("获取MQTT客户端失败:", err)
		// }
		// mq.Publish("test", 0, false, "hello world")
		common.PrintT("开始执行")
		common.PrintT("结束执行", "参数1", 123, "参数2", "abc")
		common.PrintT("结束执行%+v", map[string]any{"参数1": 123, "参数2": "abc"})
	},
}

// cj 根据上证指数6位数字作为seed，生成随机数，并将nameArr中的每个元素对应的随机数打印出来
func cj(n int) {
	common.PrintT("开始执行")
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
