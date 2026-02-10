package common

import (
	"math/rand"
	"strings"
	"time"
)

// Descrption: 字符串切割成数组
func SmartSplit(str string, sep string) []string {
	//避免空字符串返回[""]
	if str == "" {
		return []string{}
	}
	res := strings.Split(str, sep)
	return res
}

// Description:生成区间内的随机数
//
// @param s 区间开始值
//
// @param e 区间结束值
func RandInt(s int, e int) int {
	// 使用当前时间的纳秒数设置种子
	rand.New(rand.NewSource(time.Now().UnixNano()))
	// 生成一个0到99之间的随机整数
	num := rand.Intn(e)
	return num - s
}

// 日期转时间戳秒
func Date2Timestamp(dateStr string) (int64, error) {
	// 定义日期字符串的布局格式
	layout := "2006-01-02 15:04:05" // 这是 Go 中标准的时间布局格式

	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return 0, err
	}

	// 解析时间为本地时区
	t, err := time.ParseInLocation(layout, dateStr, loc)
	if err != nil {
		return 0, err
	}

	// 转换为 Unix 时间戳（秒）
	timestamp := t.Unix()
	return timestamp, nil
}
